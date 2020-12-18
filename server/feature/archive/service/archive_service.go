package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ruinnel/giregi.rip-server/common"
	"github.com/ruinnel/giregi.rip-server/domain"
	"github.com/ruinnel/giregi.rip-server/parser"
	"github.com/ruinnel/giregi.rip-server/queue"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"
)

type archiveService struct {
	archiveRepository domain.ArchiveRepository
	archiveCache      domain.ArchiveCache
	siteRepository    domain.SiteRepository
	webPageRepository domain.WebPageRepository
	tagRepository     domain.TagRepository
	cacheTimeout      time.Duration
	checkTerm         time.Duration
	queue             queue.Queue
}

type ArchiveServiceOption struct {
	CacheTimeout *time.Duration
	CheckTerm    *time.Duration
}

func NewArchiveService(
	archiveRepository domain.ArchiveRepository,
	archiveCache domain.ArchiveCache,
	siteRepository domain.SiteRepository,
	webPageRepository domain.WebPageRepository,
	tagRepository domain.TagRepository,
	q queue.Queue,
	options ...ArchiveServiceOption,
) domain.ArchiveService {
	cacheTimeout := 30 * time.Minute
	checkTerm := 30 * time.Second
	if len(options) > 0 {
		option := options[0]
		if option.CacheTimeout != nil {
			cacheTimeout = *option.CacheTimeout
		}
		if option.CheckTerm != nil {
			checkTerm = *option.CheckTerm
		}
	}
	return &archiveService{
		archiveRepository: archiveRepository,
		archiveCache:      archiveCache,
		siteRepository:    siteRepository,
		webPageRepository: webPageRepository,
		tagRepository:     tagRepository,
		cacheTimeout:      cacheTimeout,
		checkTerm:         checkTerm,
		queue:             q,
	}
}

type archiveStatus struct {
	Available     bool     `json:"available"`
	JobId         string   `json:"job_id"`
	Message       string   `json:"message"`
	OriginalJobId string   `json:"original_job_id"`
	OriginalUrl   string   `json:"original_url"`
	Resources     []string `json:"resources"`
	SecondsAgo    int      `json:"seconds_ago"`
	Status        string   `json:"status"`
	Timestamp     string   `json:"timestamp"`
}

func (s *archiveService) Archive(ctx context.Context, userId int64, targetUrl *url.URL, tags []domain.Tag, memo string, title string, public bool) (*domain.Archive, error) {
	logger := common.GetLogger()
	strippedUrl := parser.StripUrl(targetUrl)

	archive, err := s.archiveCache.Get(ctx, strippedUrl.String())
	if err != nil {
		logger.Println(fmt.Sprintf("get cache fail.(url - %s)", strippedUrl.String()))
	}

	if archive == nil {
		parsed, err := parser.Parse(strippedUrl)
		if err != nil {
			logger.Printf("archive: parse fail - %v", err)
		}
		archive = &domain.Archive{
			Summary: parsed.ToList(),
		}
	}

	site, err := s.siteRepository.GetByDomain(ctx, strippedUrl.Host)
	if site == nil {
		site = &domain.Site{Domain: strippedUrl.Host}
		err = s.siteRepository.Store(ctx, site)
		if err != nil {
			return nil, err
		}
	}

	webPage, err := s.webPageRepository.GetByURL(ctx, strippedUrl.String())
	if webPage == nil {
		webPage = &domain.WebPage{SiteID: site.ID, URL: strippedUrl.String(), Site: site}
		err = s.webPageRepository.Store(ctx, webPage)
		if err != nil {
			return nil, err
		}
	}

	webPage.Site = site
	archive.UserID = userId
	archive.WebPageID = webPage.ID
	archive.WebPage = webPage
	archive.Status = "pending"
	archive.Memo = memo
	archive.Title = title
	archive.Public = public

	err = s.archiveRepository.Store(ctx, archive)
	if err != nil {
		return nil, err
	}

	for _, tag := range tags {
		if tag.ID < 0 {
			err := s.tagRepository.Store(ctx, &tag)
			if err != nil {
				logger.Printf("archive: store tag - %v", err)
				return nil, err
			}
		}
		exists, err := s.tagRepository.ExistsMapping(ctx, archive.ID, tag.ID)
		if err != nil {
			logger.Printf("archive: exist mapping - %v", err)
			return nil, err
		}
		logger.Printf("archive: exists mapping - %v", exists)

		if !exists {
			err = s.tagRepository.AddMapping(ctx, archive.ID, tag.ID)
			if err != nil {
				logger.Printf("archive: add mapping - %v", err)
				return nil, err
			}
		}
	}

	return s.ProcessArchive(ctx, archive)
}

func (s *archiveService) GetByURL(ctx context.Context, userId int64, targetUrl *url.URL) (*domain.Archive, error) {
	strippedUrl := parser.StripUrl(targetUrl)
	webPage, err := s.getWebPageByUrl(ctx, strippedUrl.String())
	if err != nil {
		return nil, err
	}

	if webPage == nil {
		return nil, nil
	}

	conditions := []common.Condition{
		{Field: domain.ArchiveField.UserID, Op: common.Eq, Val: userId},
		{Field: domain.ArchiveField.WebPageID, Op: common.Eq, Val: webPage.ID},
	}
	return s.archiveRepository.One(ctx, conditions)
}

func (s *archiveService) Preview(ctx context.Context, userId int64, targetUrl *url.URL) (*domain.Archive, error) {
	logger := common.GetLogger()
	strippedUrl := parser.StripUrl(targetUrl)
	parsed, err := parser.Parse(strippedUrl)
	if err != nil {
		logger.Printf("archive: parse fail - %v", err)
	}

	site, err := s.siteRepository.GetByDomain(ctx, strippedUrl.Host)
	if site == nil {
		site = &domain.Site{Domain: strippedUrl.Host}
	}

	webPage, err := s.webPageRepository.GetByURL(ctx, strippedUrl.String())
	if webPage == nil {
		webPage = &domain.WebPage{URL: strippedUrl.String(), Site: site}
	}

	archive := &domain.Archive{
		WebPage: webPage,
		Summary: parsed.ToList(),
	}

	err = s.archiveCache.Set(ctx, archive, s.cacheTimeout)
	if err != nil {
		logger.Println("cache set fail.", err)
	}
	return archive, nil
}

func (s *archiveService) ProcessArchive(ctx context.Context, archive *domain.Archive) (*domain.Archive, error) {
	logger := common.GetLogger()
	result, err := s.processArchive(ctx, archive)
	if err != nil {
		logger.Printf("archive process fail: %v", err)
		return nil, err
	} else {
		err = s.archiveRepository.Update(ctx, archive)
		if err != nil {
			logger.Printf("archive update fail: %v", err)
			return nil, err
		}
	}

	err = s.queue.Enqueue(archive)
	if err != nil {
		logger.Printf("request check fail: %v", err)
		return nil, err
	}
	return result, nil
}

func (s *archiveService) Fetch(ctx context.Context, params domain.ArchiveFetchParams, cursor string, count int) (*common.FetchResult, error) {
	return s.archiveRepository.Fetch(ctx, params, cursor, count)
}

func (s *archiveService) GetByID(ctx context.Context, id int64) (*domain.Archive, error) {
	return s.archiveRepository.GetByID(ctx, id)
}

func (s *archiveService) getWebPageByUrl(ctx context.Context, url string) (*domain.WebPage, error) {
	conditions := []common.Condition{
		{
			Field: domain.WebPageField.URL,
			Op:    common.Eq,
			Val:   url,
		},
	}
	exists, err := s.webPageRepository.Exists(ctx, conditions)
	if exists {
		return s.webPageRepository.One(ctx, conditions)
	} else {
		return nil, err
	}
}

func (s *archiveService) processArchive(ctx context.Context, archive *domain.Archive) (*domain.Archive, error) {
	pageUrl := archive.WebPage.URL
	logger := common.GetLogger()
	requestUrl := fmt.Sprintf("http://web.archive.org/save/%s", pageUrl)

	res, err := http.PostForm(requestUrl, url.Values{"url": {pageUrl}})
	if res == nil {
		logger.Printf("archive fail(response is nil)")
		return nil, errors.New("archive fail(response is nil)")
	}

	errorHeader := res.Header.Get("X-Archive-Wayback-Runtime-Error")
	if errorHeader != "" {
		logger.Printf("archive fail(Runtime-Error): %s", errorHeader)
		return nil, errors.New(fmt.Sprintf("archive fail(Runtime-Error): %s", errorHeader))
	}

	liveWebError := res.Header.Get("x-archive-wayback-liveweb-error")
	if liveWebError != "" {
		logger.Printf("archive fail(Wayback-LiveWeb-Error): %s", errorHeader)
		return nil, errors.New(fmt.Sprintf("archive fail(Wayback-LiveWeb-Error): %s", errorHeader))
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logger.Printf("archive fail(read body): %v", err)
		return nil, err
	}

	// fmt.Println("body - ", string(body))
	strings.Split(string(body), "\n")
	// spn.watchJob("7ed1c1b7-e6d9-40ef-a36d-f11268f5bfab", "/_static/",
	regex := regexp.MustCompile(`(?s)spn.watchJob\("([0-9a-fA-F\-]+)"`)
	// regex := regexp.MustCompile(`(?s)var JOB_ID = "([0-9a-fA-F\-]+)";(?s)`)
	match := regex.FindStringSubmatch(string(body))
	if len(match) >= 2 {
		jobId := match[1]
		archive.JobID = jobId
		return archive, nil
	} else {
		logger.Printf("archive fail(jobId not found)")
		return nil, errors.New("archive fail(jobId not found)")
	}
}

func (s *archiveService) CheckProgress(ctx context.Context, archive *domain.Archive) error {
	logger := common.GetLogger()
	status, err := s.checkStatus(archive.JobID)
	if err != nil {
		logger.Printf("archive check progress fail: %v", err)
		return err
	}

	if status.Status == "success" {
		logger.Printf("archive check success.")
		archive.Status = status.Status
		archive.WaybackID = fmt.Sprintf("/web/%s/%s", status.Timestamp, status.OriginalUrl)
		return s.archiveRepository.Update(ctx, archive)
	} else if status.Status == "pending" {
		logger.Printf(fmt.Sprintf("archive check progress retry... after %v second.", time.Duration(s.checkTerm).Seconds()))
		time.Sleep(s.checkTerm)
		return s.CheckProgress(ctx, archive)
	} else {
		return errors.New(fmt.Sprintf("archive check - unknown status (%s)", status.Status))
	}
}

func (s *archiveService) checkStatus(jobId string) (*archiveStatus, error) {
	logger := common.GetLogger()
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)
	requestUrl := fmt.Sprintf("http://web.archive.org/save/status/%s?_t=%d", jobId, timestamp)

	res, err := http.Get(requestUrl)
	if res == nil {
		logger.Printf("archive check status fail(response is nil)")
		return nil, errors.New("archive check status fail(response is nil)")
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logger.Printf("archive check status fail(read body): %v", err)
		return nil, err
	}

	status := new(archiveStatus)
	err = json.Unmarshal(body, status)
	if err != nil {
		logger.Printf("archive check status fail(marshal json): %v", err)
		return nil, err
	}
	return status, nil
}
