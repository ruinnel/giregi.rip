package domain

import (
	"context"
	"github.com/ruinnel/giregi.rip-server/common"
	"reflect"
)

type WebPage struct {
	ID        int64   `json:"id" mysql:"id" storm:"id,increment"`
	SiteID    int64   `json:"siteId" mysql:"site_id" storm:"index"`
	URL       string  `json:"url" mysql:"url" storm:"index"`
	Title     *string `json:"title,omitempty" mysql:"title" storm:"index"`
	CreatedAt Time    `json:"createdAt" mysql:"created_at" storm:"index"`
	UpdatedAt Time    `json:"updatedAt" mysql:"updated_at"`

	Site *Site `json:"site"`
}

type webPageField struct {
	ID        reflect.StructField
	SiteID    reflect.StructField
	URL       reflect.StructField
	Title     reflect.StructField
	CreatedAt reflect.StructField
	UpdatedAt reflect.StructField
}

var WebPageField = makeFields(&WebPage{}, &webPageField{}).(*webPageField)

type WebPageRepository interface {
	GetByURL(ctx context.Context, url string) (*WebPage, error)
	Fetch(ctx context.Context, conditions []common.Condition, cursor string, count int) (data []WebPage, nextCursor string, err error)
	One(ctx context.Context, conditions []common.Condition) (*WebPage, error)
	Exists(ctx context.Context, conditions []common.Condition) (bool, error)
	GetByID(ctx context.Context, id int64) (*WebPage, error)
	Store(ctx context.Context, user *WebPage) error
	Update(ctx context.Context, user *WebPage) error
	Delete(ctx context.Context, id int64) error
}
