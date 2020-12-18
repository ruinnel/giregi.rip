package domain

import (
	"context"
	"github.com/ruinnel/giregi.rip-server/common"
	"reflect"
)

type Tag struct {
	ID        int64  `json:"id" mysql:"id" storm:"id,increment"`
	Name      string `json:"name" mysql:"name" storm:"index"`
	UserID    int64  `json:"userId" mysql:"user_id" storm:"index"`
	CreatedAt Time   `json:"createdAt" mysql:"created_at" storm:"index"`
	UpdatedAt Time   `json:"updatedAt" mysql:"updated_at"`
}

type tagField struct {
	ID        reflect.StructField
	Name      reflect.StructField
	UserID    reflect.StructField
	CreatedAt reflect.StructField
	UpdatedAt reflect.StructField
}

var TagField = makeFields(&Tag{}, &tagField{}).(*tagField)

type TagRepository interface {
	Fetch(ctx context.Context, conditions []common.Condition, cursor string, count int) (*common.FetchResult, error)
	One(ctx context.Context, conditions []common.Condition) (*Tag, error)
	Exists(ctx context.Context, conditions []common.Condition) (bool, error)
	GetByID(ctx context.Context, id int64) (*Tag, error)
	Store(ctx context.Context, tag *Tag) error
	Update(ctx context.Context, tag *Tag) error
	Delete(ctx context.Context, id int64) error
	ExistsMapping(ctx context.Context, archiveId int64, tagId int64) (bool, error)
	AddMapping(ctx context.Context, archiveId int64, tagId int64) error
}
