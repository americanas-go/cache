package bigcache

import (
	"context"

	"github.com/allegro/bigcache/v3"
	"github.com/americanas-go/cache"
)

type driver struct {
	cache *bigcache.BigCache
}

func (c *driver) Del(ctx context.Context, key string) error {
	return c.cache.Delete(key)
}

func (c *driver) Get(ctx context.Context, key string) (data []byte, err error) {
	return c.cache.Get(key)
}

func (c *driver) Set(ctx context.Context, key string, data []byte) (err error) {

	if err = c.cache.Set(key, data); err != nil {
		return err
	}

	return nil
}

func New(cache *bigcache.BigCache) cache.Driver {
	return &driver{cache: cache}
}
