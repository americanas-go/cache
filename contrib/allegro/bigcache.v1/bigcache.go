package bigcache

import (
	"context"

	"github.com/allegro/bigcache"
	"github.com/americanas-go/cache"
)

type bcache struct {
	cache *bigcache.BigCache
}

func (c *bcache) Del(ctx context.Context, key string) error {
	return c.cache.Delete(key)
}

func (c *bcache) Get(ctx context.Context, key string) (data []byte, err error) {
	return c.cache.Get(key)
}

func (c *bcache) Set(ctx context.Context, key string, data []byte) (err error) {

	if err = c.cache.Set(key, data); err != nil {
		return err
	}

	return nil
}

func New(cache *bigcache.BigCache) cache.Driver {
	return &bcache{cache: cache}
}
