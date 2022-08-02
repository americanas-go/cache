package redis

import (
	"context"

	"github.com/americanas-go/cache"
	"github.com/go-redis/redis/v7"
)

type clusterDriver struct {
	cache   *redis.ClusterClient
	options *Options
}

func (c *clusterDriver) Del(ctx context.Context, key string) error {
	return c.cache.WithContext(ctx).Del(key).Err()
}

func (c *clusterDriver) Get(ctx context.Context, key string) (data []byte, err error) {
	return c.cache.WithContext(ctx).Get(key).Bytes()
}

func (c *clusterDriver) Set(ctx context.Context, key string, data []byte) (err error) {

	c.cache.WithContext(ctx).Del(key)

	if err = c.cache.WithContext(ctx).Set(key, data, c.options.TTL).Err(); err != nil {
		return err
	}

	return nil
}

func NewCluster(cache *redis.ClusterClient, options *Options) cache.Driver {
	return &clusterDriver{cache: cache, options: options}
}
