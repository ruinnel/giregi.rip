package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ruinnel/giregi.rip-server/common"
	"github.com/ruinnel/giregi.rip-server/domain"
	"github.com/ruinnel/giregi.rip-server/parser"
	"github.com/streadway/amqp"
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
	rabbitMQ          common.RabbitMQ
	cacheTimeout      time.Duration
	checkTerm         time.Duration
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
	rabbitMQ common.RabbitMQ,
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
		rabbitMQ:          rabbitMQ,
		cacheTimeout:      cacheTimeout,
		checkTerm:         checkTerm,
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

func (s *archiveService) Archive(ctx context.Context, userId int64, targetUrl *url.URL, tags []domain.Tag, memo string, public bool) (*domain.Archive, error) {
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

	err = s.RequestArchive(ctx, archive)
	if err != nil {
		return nil, err
	}
	return archive, nil
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
		{Field: "user_id", Op: common.Eq, Val: userId},
		{Field: "web_page_id", Op: common.Eq, Val: webPage.ID},
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

func (s *archiveService) RequestArchive(ctx context.Context, archive *domain.Archive) error {
	logger := common.GetLogger()
	rabbitMQUrl := fmt.Sprintf("amqp://%s:%s@%s:%d/", s.rabbitMQ.Username, s.rabbitMQ.Password, s.rabbitMQ.Host, s.rabbitMQ.Port)
	fmt.Print(rabbitMQUrl)
	conn, err := amqp.Dial(rabbitMQUrl)
	if err != nil {
		logger.Printf("amqp - %v", rabbitMQUrl)
		logger.Printf("connect to rabbitMQ(amqp) fail: %v", err)
		return err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		logger.Printf("connect to rabbitMQ(amqp) fail: %v", err)
		return err
	}
	defer ch.Close()

	_, err = ch.QueueDeclare(s.rabbitMQ.Queue, false, false, false, false, nil)
	if err != nil {
		logger.Printf("connect to rabbitMQ(amqp) fail: %v", err)
		return err
	}

	data, err := json.Marshal(archive)
	if err != nil {
		logger.Printf("Publish: marshal message fail: %v", err)
		return err
	}

	err = ch.Publish("", s.rabbitMQ.Queue, false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        data,
	})

	if err != nil {
		logger.Printf("Publish message fail: %v", err)
		return err
	}
	return nil
}

func (s *archiveService) ProcessArchive(ctx context.Context, archive *domain.Archive) error {
	return s.processArchive(archive)
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
			Field: "url",
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

func (s *archiveService) processArchive(archive *domain.Archive) error {
	pageUrl := archive.WebPage.URL
	logger := common.GetLogger()
	requestUrl := fmt.Sprintf("http://web.archive.org/save/%s", pageUrl)

	res, err := http.PostForm(requestUrl, url.Values{"url": {pageUrl}})
	if res == nil {
		logger.Printf("archive fail(response is nil)")
		return errors.New("archive fail(response is nil)")
	}

	errorHeader := res.Header.Get("X-Archive-Wayback-Runtime-Error")
	if errorHeader != "" {
		logger.Printf("archive fail(Runtime-Error): %s", errorHeader)
		return errors.New(fmt.Sprintf("archive fail(Runtime-Error): %s", errorHeader))
	}

	liveWebError := res.Header.Get("x-archive-wayback-liveweb-error")
	if liveWebError != "" {
		logger.Printf("archive fail(Wayback-LiveWeb-Error): %s", errorHeader)
		return errors.New(fmt.Sprintf("archive fail(Wayback-LiveWeb-Error): %s", errorHeader))
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logger.Printf("archive fail(read body): %v", err)
		return err
	}

	// fmt.Println("body - ", string(body))
	strings.Split(string(body), "\n")
	// spn.watchJob("7ed1c1b7-e6d9-40ef-a36d-f11268f5bfab", "/_static/",
	regex := regexp.MustCompile(`(?s)spn.watchJob\("([0-9a-fA-F\-]+)"`)
	// regex := regexp.MustCompile(`(?s)var JOB_ID = "([0-9a-fA-F\-]+)";(?s)`)
	match := regex.FindStringSubmatch(string(body))
	if len(match) >= 2 {
		jobId := match[1]
		status := s.checkProgress(jobId)
		archive.Status = status.Status
		archive.JobID = jobId
		archive.WaybackID = fmt.Sprintf("/web/%s/%s", status.Timestamp, status.OriginalUrl)
		return s.archiveRepository.Update(context.Background(), archive)
	} else {
		logger.Printf("archive fail(jobId not found)")
		return errors.New("archive fail(jobId not found)")
	}
}

func (s *archiveService) checkProgress(jobId string) *archiveStatus {
	logger := common.GetLogger()
	status, err := s.checkStatus(jobId)
	if err != nil {
		logger.Printf("archive check progress fail: %v", err)
		return nil
	}

	if status.Status == "success" {
		logger.Printf("archive check success.")
		return status
	} else if status.Status == "pending" {
		logger.Printf("archive check progress retry... after 10 second.")
		time.Sleep(s.checkTerm)
		return s.checkProgress(jobId)
	} else {
		return nil
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