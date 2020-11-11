package domain

import (
	"context"
	"github.com/ruinnel/giregi.rip-server/common"
)

type Site struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Domain    string `json:"domain"`
	CreatedAt Time   `json:"createdAt"`
	UpdatedAt Time   `json:"updatedAt"`

	WebPages []WebPage `json:"webPages"`
	Tags     []Tag     `json:"tags"`
}

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
