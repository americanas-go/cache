package freecache

import (
	"context"

	"github.com/americanas-go/cache"
	"github.com/go-redis/redis/v8"
)

type cluster struct {
	cache   *redis.ClusterClient
	options *Options
}

func (c *cluster) Del(ctx context.Context, key string) error {
	return c.cache.WithContext(ctx).Del(ctx, key).Err()
}

func (c *cluster) Get(ctx context.Context, key string) (data []byte, err error) {
	return c.cache.WithContext(ctx).Get(ctx, key).Bytes()
}

func (c *cluster) Set(ctx context.Context, key string, data []byte) (err error) {

	c.cache.WithContext(ctx).Del(ctx, key)

	if err = c.cache.WithContext(ctx).Set(ctx, key, data, c.options.TTL).Err(); err != nil {
		return err
	}

	return nil
}

func NewCluster(cache *redis.ClusterClient, options *Options) cache.Driver {
	return &cluster{cache: cache, options: options}
}
