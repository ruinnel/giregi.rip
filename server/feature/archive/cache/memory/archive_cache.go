package memory

import (
	"context"
	"github.com/patrickmn/go-cache"
	"github.com/ruinnel/giregi.rip-server/domain"
	"time"
)

type archiveCache struct {
	Cache *cache.Cache
}

func NewArchiveCache(cache *cache.Cache) domain.ArchiveCache {
	return &archiveCache{Cache: cache}
}

func (r archiveCache) Get(ctx context.Context, url string) (*domain.Archive, error) {
	val, exists := r.Cache.Get(url)
	if exists {
		return val.(*domain.Archive), nil
	} else {
		return nil, nil
	}
}

func (r archiveCache) Set(ctx context.Context, archive *domain.Archive, duration time.Duration) error {
	url := archive.WebPage.URL
	r.Cache.Set(url, archive, duration)
	return nil
}
