package domain

import (
	"context"
	"github.com/ruinnel/giregi.rip-server/common"
)

type Tag struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	UserID    int64  `json:"userId"`
	CreatedAt Time   `json:"createdAt"`
	UpdatedAt Time   `json:"updatedAt"`
}

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
