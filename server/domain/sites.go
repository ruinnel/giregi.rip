package domain

import (
	"context"
	"github.com/ruinnel/giregi.rip-server/common"
	"reflect"
)

type Site struct {
	ID        int64  `json:"id" mysql:"id" storm:"id,increment"`
	Name      string `json:"name" mysql:"name" storm:"index"`
	Domain    string `json:"domain" mysql:"domain" storm:"index"`
	CreatedAt Time   `json:"createdAt" mysql:"created_at" storm:"index"`
	UpdatedAt Time   `json:"updatedAt" mysql:"updated_at"`

	WebPages []WebPage `json:"webPages"`
	Tags     []Tag     `json:"tags"`
}

type siteField struct {
	ID        reflect.StructField
	Name      reflect.StructField
	Domain    reflect.StructField
	CreatedAt reflect.StructField
	UpdatedAt reflect.StructField
}

var SiteField = makeFields(&Site{}, &siteField{}).(*siteField)

type SiteRepository interface {
	GetByDomain(ctx context.Context, domain string) (*Site, error)
	Fetch(ctx context.Context, conditions []common.Condition, cursor string, count int) (data []Site, nextCursor string, err error)
	One(ctx context.Context, conditions []common.Condition) (*Site, error)
	Exists(ctx context.Context, conditions []common.Condition) (bool, error)
	GetByID(ctx context.Context, id int64) (*Site, error)
	Store(ctx context.Context, user *Site) error
	Update(ctx context.Context, user *Site) error
	Delete(ctx context.Context, id int64) error
}
