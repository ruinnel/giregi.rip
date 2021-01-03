package redis

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"github.com/ruinnel/giregi.rip-server/domain"
	"time"
)

type archiveCache struct {
	Client *redis.Client
}

func (r archiveCache) Get(ctx context.Context, url string) (*domain.Archive, error) {
	val, err := r.Client.Get(ctx, url).Result()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	archive := new(domain.Archive)
	err = json.Unmarshal([]byte(val), archive)
	if err != nil {
		return nil, err
	}
	return archive, nil
}

func (r archiveCache) Set(ctx context.Context, archive *domain.Archive, duration time.Duration) error {
	url := archive.WebPage.URL
	val, err := json.Marshal(archive)
	if err != nil {
		return err
	}
	return r.Client.SetNX(ctx, url, string(val), duration).Err()
}

func NewArchiveCache(client *redis.Client) domain.ArchiveCache {
	return &archiveCache{Client: client}
}
