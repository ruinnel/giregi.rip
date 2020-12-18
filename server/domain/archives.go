package domain

import (
	"context"
	"github.com/ruinnel/giregi.rip-server/common"
	"net/url"
	"reflect"
	"time"
)

type Archive struct {
	ID        int64                    `json:"id" mysql:"id" storm:"id,increment"`
	WebPageID int64                    `json:"webPageId" mysql:"web_page_id" storm:"index"`
	UserID    int64                    `json:"userId" mysql:"user_id" storm:"index"`
	Memo      string                   `json:"memo" mysql:"memo"`
	Title     string                   `json:"title" mysql:"title" storm:"index"`
	Status    string                   `json:"status" mysql:"status"`
	JobID     string                   `json:"jobId" mysql:"job_id"`
	WaybackID string                   `json:"waybackId,omitempty" mysql:"wayback_id"  storm:"index"`
	Summary   []map[string]interface{} `json:"summary" mysql:"summary"`
	Public    bool                     `json:"public" mysql:"public"`
	CreatedAt Time                     `json:"createdAt" mysql:"created_at" storm:"index"`
	UpdatedAt Time                     `json:"updatedAt" mysql:"updated_at"`

	TagIDs []int64 `json:"tagIds" storm:"index"`

	Tags    []Tag    `json:"tags"`
	WebPage *WebPage `json:"webPage"`
	User    *User    `json:"user"`
}

type ArchiveFetchParams struct {
	UserID  int64
	TagID   int64
	Keyword string
}

type archiveField struct {
	ID        reflect.StructField
	WebPageID reflect.StructField
	UserID    reflect.StructField
	Memo      reflect.StructField
	Title     reflect.StructField
	Status    reflect.StructField
	JobID     reflect.StructField
	WaybackID reflect.StructField
	Summary   reflect.StructField
	Public    reflect.StructField
	CreatedAt reflect.StructField
	UpdatedAt reflect.StructField
	TagIDs    reflect.StructField
}

var ArchiveField = makeFields(&Archive{}, &archiveField{}).(*archiveField)

type ArchiveService interface {
	Archive(ctx context.Context, userId int64, targetUrl *url.URL, tags []Tag, memo string, title string, public bool) (*Archive, error)
	GetByURL(ctx context.Context, userId int64, targetUrl *url.URL) (*Archive, error)
	Preview(ctx context.Context, userId int64, targetUrl *url.URL) (*Archive, error)
	ProcessArchive(ctx context.Context, archive *Archive) (*Archive, error)
	CheckProgress(ctx context.Context, archive *Archive) error
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
