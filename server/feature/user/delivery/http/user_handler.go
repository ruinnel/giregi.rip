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
)

func login(userService domain.UserService) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		logger := common.GetRequestLogger()
		ctx := request.Context()
		param := new(struct {
			Email   string `schema:"email" validate:"required,email"`
			IdToken string `schema:"idToken" validate:"required"`
			TokenId string `schema:"tokenId"`
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
			common.WriteError(writer, common.NewInvalidParamError(fmt.Sprintf("invalid param: %v", param)), err)
			return
		}

		tokenId, err := common.DecodeHashId(param.TokenId)
		if err != nil {
			logger.Printf(request, "invalid tokenId: %v", param.TokenId)
			common.WriteError(writer, common.NewInvalidParamError(fmt.Sprintf("invalid tokenId: %v", param.TokenId)), err)
			return
		}

		userAgent := request.Header.Get("User-Agent")

		token, err := userService.Login(ctx, param.Email, param.IdToken, tokenId, userAgent)
		if err != nil {
			logger.Printf(request, "login fail: %v", err)
			common.WriteError(writer, common.NewInvalidParamError(fmt.Sprintf("login fail: %v", param)), err)
			return
		}

		common.WriteJson(writer, token)
	}
}

func logout(userService domain.UserService) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		logger := common.GetRequestLogger()
		ctx := request.Context()
		accessToken := request.Header.Get("Authorization")
		err := userService.Logout(ctx, accessToken)
		if err != nil {
			logger.Printf(request, "logout fail: %v", err)
			common.WriteError(writer, common.NewUnknownError(fmt.Sprintf("logout fail: %v", err), err), err)
			return
		}
		common.WriteJson(writer, "OK")
	}
}

func my(userService domain.UserService) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		logger := common.GetRequestLogger()
		user := middleware.GetUser(request)
		if user == nil {
			logger.Printf(request, "invalid session")
			common.WriteError(writer, common.NewInvalidParamError("invalid session"), nil)
			return
		}
		common.WriteJson(writer, user)
	}
}

func profile(userService domain.UserService) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		logger := common.GetRequestLogger()
		ctx := request.Context()
		param := new(struct {
			Id int64 `schema:"id" validate:"omitempty,alphanum"`
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

		user, err := userService.GetByID(ctx, param.Id)
		if err != nil {
			logger.Printf(request, "user id not registered: %v", param.Id)
			common.WriteError(writer, common.NewInvalidParamError(fmt.Sprintf("user id not registered: %v", param.Id)), err)
			return
		}
		common.WriteJson(writer, user)
	}
}

func tags(userService domain.UserService) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		logger := common.GetRequestLogger()
		ctx := request.Context()

		tags, err := userService.Tags(ctx)
		if err != nil {
			logger.Printf(request, "get user tags fail")
			common.WriteError(writer, common.NewInvalidParamError("get user tags fail"), err)
			return
		}
		common.WriteJson(writer, tags)
	}
}

func archives(userService domain.UserService) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		logger := common.GetRequestLogger()
		ctx := request.Context()
		user := middleware.GetUser(request)
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
			logger.Printf(request, "invalid param: %v - %v", param, err)
			common.WriteError(writer, common.NewInvalidParamError(fmt.Sprintf("invalid param: %v", param), err), err)
			return
		}

		if param.Count == 0 {
			param.Count = common.DefaultPagingCount
		}

		fetchParams := domain.ArchiveFetchParams{
			UserID:  user.ID,
			Keyword: param.Keyword,
			TagID:   param.TagID,
		}
		result, err := userService.GetArchives(ctx, fetchParams, param.NextCursor, param.Count)
		if err != nil {
			logger.Printf(request, "get user archives fail: (%v), %v", param.Keyword, err)
			common.WriteError(writer, common.NewInvalidParamError(fmt.Sprintf("get user archives fail: %v", param.Keyword), err), err)
			return
		}

		common.WriteJson(writer, result)
	}
}

func User(router *mux.Router, userService domain.UserService) {
	router.HandleFunc("", my(userService)).Methods(http.MethodGet)
	router.HandleFunc("/{id:[0-9]+}", profile(userService)).Methods(http.MethodGet)
	router.HandleFunc("/login", login(userService)).Methods(http.MethodPost)
	router.HandleFunc("/logout", logout(userService)).Methods(http.MethodDelete)
	router.HandleFunc("/tags", tags(userService)).Methods(http.MethodGet)
	router.HandleFunc("/archives", archives(userService)).Methods(http.MethodGet)
}
