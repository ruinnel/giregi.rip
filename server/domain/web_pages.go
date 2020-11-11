package domain

import (
	"context"
	"github.com/ruinnel/giregi.rip-server/common"
)

type WebPage struct {
	ID        int64   `json:"id"`
	SiteID    int64   `json:"siteId"`
	URL       string  `json:"url"`
	Title     *string `json:"title,omitempty"`
	CreatedAt Time    `json:"createdAt"`
	UpdatedAt Time    `json:"updatedAt"`

	Site *Site `json:"site"`
}

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
