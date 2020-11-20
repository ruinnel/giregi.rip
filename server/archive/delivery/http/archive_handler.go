package http

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/ruinnel/giregi.rip-server/common"
	"github.com/ruinnel/giregi.rip-server/domain"
	"github.com/ruinnel/giregi.rip-server/http/middleware"
	"net/http"
	"net/url"
)

func archive(archiveService domain.ArchiveService) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		logger := common.GetRequestLogger()
		ctx := request.Context()
		user := middleware.GetUser(request)
		param := new(struct {
			Url    string `json:"url" schema:"url" validate:"required,url"`
			Memo   string `json:"memo" schema:"memo" validate:"required,min=1,max=200"`
			Title  string `json:"title" schema:"title" validate:"required,min=1,max=200"`
			Public bool   `json:"public" schema:"public" validate:"omitempty"`
			Tags   []struct {
				ID   int64  `json:"id" schema:"id"`
				Name string `json:"name" schema:"name" validate:"required"`
			} `json:"tags" schema:"tags"`
		})
		err := common.ParseJsonBody(request, param)
		if err != nil {
			logger.Printf(request, "parse json body fail: %v", err)
			common.WriteError(writer, common.NewInvalidParamError("parse json body fail", err), err)
			return
		}
		validate := validator.New()
		err = validate.Struct(param)
		if err != nil {
			logger.Printf(request, "invalid param: %v", param)
			common.WriteError(writer, common.NewInvalidParamError(fmt.Sprintf("invalid param: %v", param), err), err)
			return
		}
		targetUrl, err := url.Parse(param.Url)
		if err != nil {
			logger.Printf(request, "invalid url: %v", param.Url)
			common.WriteError(writer, common.NewInvalidParamError(fmt.Sprintf("invalid url: %v", param.Url), err), err)
			return
		}

		tags := make([]domain.Tag, len(param.Tags))
		for idx, tag := range param.Tags {
			tags[idx] = domain.Tag{ID: tag.ID, UserID: user.ID, Name: tag.Name}
		}
		logger.Printf(request, "tags - %v", tags)
		archive, err := archiveService.Archive(ctx, user.ID, targetUrl, tags, param.Memo, param.Title, param.Public)
		if err != nil {
			logger.Printf(request, "archive fail: (%v), %v", param.Url, err)
			common.WriteError(writer, common.NewInvalidParamError(fmt.Sprintf("archive fail: %v", param.Url), err), err)
			return
		}

		common.WriteJson(writer, archive)
	}
}

func getByUrl(archiveService domain.ArchiveService) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		logger := common.GetRequestLogger()
		ctx := request.Context()
		user := middleware.GetUser(request)
		param := new(struct {
			Url string `schema:"url" validate:"required,url"`
		})
		decoder := schema.NewDecoder()
		_ = decoder.Decode(param, request.URL.Query())
		validate := validator.New()
		err := validate.Struct(param)
		if err != nil {
			logger.Printf(request, "invalid param: %v", param)
			common.WriteError(writer, common.NewInvalidParamError(fmt.Sprintf("invalid param: %v", param), err), err)
			return
		}
		targetUrl, err := url.Parse(param.Url)
		if err != nil {
			logger.Printf(request, "invalid url: %v", param.Url)
			common.WriteError(writer, common.NewInvalidParamError(fmt.Sprintf("invalid url: %v", param.Url), err), err)
			return
		}

		archive, err := archiveService.GetByURL(ctx, user.ID, targetUrl)
		if err != nil {
			logger.Printf(request, "getByUrl fail: (%v), %v", param.Url, err)
			common.WriteError(writer, common.NewInvalidParamError(fmt.Sprintf("getByUrl fail: %v", param.Url), err), err)
			return
		}

		common.WriteJson(writer, archive)
	}
}

func preview(archiveService domain.ArchiveService) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		logger := common.GetRequestLogger()
		ctx := request.Context()
		user := middleware.GetUser(request)
		param := new(struct {
			Url string `schema:"url" validate:"required,url"`
		})
		decoder := schema.NewDecoder()
		_ = decoder.Decode(param, request.URL.Query())
		validate := validator.New()
		err := validate.Struct(param)
		if err != nil {
			logger.Printf(request, "invalid param: %v", param)
			common.WriteError(writer, common.NewInvalidParamError(fmt.Sprintf("invalid param: %v", param), err), err)
			return
		}
		targetUrl, err := url.Parse(param.Url)
		if err != nil {
			logger.Printf(request, "invalid url: %v", param.Url)
			common.WriteError(writer, common.NewInvalidParamError(fmt.Sprintf("invalid url: %v", param.Url), err), err)
			return
		}

		preview, err := archiveService.Preview(ctx, user.ID, targetUrl)
		if err != nil {
			logger.Printf(request, "preview fail: %v", param.Url)
			common.WriteError(writer, common.NewInvalidParamError(fmt.Sprintf("preview fail: %v", param.Url), err), err)
			return
		}

		common.WriteJson(writer, preview)
	}
}

func getArchives(archiveService domain.ArchiveService) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		logger := common.GetRequestLogger()
		ctx := request.Context()
		// user := middleware.GetUser(request)
		param := new(struct {
			Keyword    string `schema:"keyword" validate:"omitempty,min=2"`
			TagID      int64  `schema:"tagId" validate:"omitempty,min=1"`
			NextCursor string `schema:"nextCursor,omitempty"`
			Count      int    `schema:"count" validate:"omitempty,min=1,max=100"`
		})
		decoder := schema.NewDecoder()
		_ = decoder.Decode(param, request.URL.Query())
		validate := validator.New()
		err := validate.Struct(param)
		if err != nil {
			logger.Printf(request, "invalid param: %v", param)
			common.WriteError(writer, common.NewInvalidParamError(fmt.Sprintf("invalid param: %v", param), err), err)
			return
		}

		if param.Count == 0 {
			param.Count = common.DefaultPagingCount
		}

		fetchParams := domain.ArchiveFetchParams{
			UserID:  0,
			Keyword: param.Keyword,
			TagID:   param.TagID,
		}
		result, err := archiveService.Fetch(ctx, fetchParams, param.NextCursor, param.Count)
		if err != nil {
			logger.Printf(request, "getArchives fail: (%v), %v", param.Keyword, err)
			common.WriteError(writer, common.NewInvalidParamError(fmt.Sprintf("getArchives fail: %v", param.Keyword), err), err)
			return
		}

		common.WriteJson(writer, result)
	}
}

func Archive(router *mux.Router, archiveService domain.ArchiveService) {
	router.HandleFunc("", archive(archiveService)).Methods(http.MethodPost)
	router.HandleFunc("/url", getByUrl(archiveService)).Methods(http.MethodGet)
	router.HandleFunc("/preview", preview(archiveService)).Methods(http.MethodGet)
	router.HandleFunc("", getArchives(archiveService)).Methods(http.MethodGet)
}
