package domain

import (
	"context"
	"github.com/ruinnel/giregi.rip-server/common"
	"net/url"
	"time"
)

type Archive struct {
	ID        int64                    `json:"id"`
	WebPageID int64                    `json:"webPageId"`
	UserID    int64                    `json:"userId"`
	Memo      string                   `json:"memo"`
	Status    string                   `json:"status"`
	JobID     string                   `json:"jobId"`
	WaybackID string                   `json:"waybackId,omitempty"`
	Summary   []map[string]interface{} `json:"summary"`
	Public    bool                     `json:"public"`
	CreatedAt Time                     `json:"createdAt"`
	UpdatedAt Time                     `json:"updatedAt"`

	TagIDs []int64 `json:"tagIds"`

	Tags    []Tag    `json:"tags"`
	WebPage *WebPage `json:"webPage"`
	User    *User    `json:"user"`
}

type ArchiveFetchParams struct {
	UserID  int64
	TagID   int64
	Keyword string
}

type ArchiveService interface {
	Archive(ctx context.Context, userId int64, targetUrl *url.URL, tags []Tag, memo string, public bool) (*Archive, error)
	GetByURL(ctx context.Context, userId int64, targetUrl *url.URL) (*Archive, error)
	Preview(ctx context.Context, userId int64, targetUrl *url.URL) (*Archive, error)
	RequestArchive(ctx context.Context, archive *Archive) error
	ProcessArchive(ctx context.Context, archive *Archive) error
	Fetch(ctx context.Context, params ArchiveFetchParams, cursor string, count int) (*common.FetchResult, error)
	GetByID(ctx context.Context, id int64) (*Archive, error)
}

type ArchiveRepository interface {
	Fetch(ctx context.Context, params ArchiveFetchParams, cursor string, count int) (*common.FetchResult, error)
	One(ctx context.Context, conditions []common.Condition) (*Archive, error)
	Exists(ctx context.Context, conditions []common.Condition) (bool, error)
	GetByID(ctx context.Context, id int64) (*Archive, error)
	Store(ctx context.Context, archive *Archive) error
	Update(ctx context.Context, archive *Archive) error
	Delete(ctx context.Context, id int64) error
}

type ArchiveCache interface {
	Get(ctx context.Context, url string) (*Archive, error)
	Set(ctx context.Context, archive *Archive, duration time.Duration) error
}
