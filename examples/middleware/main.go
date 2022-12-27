package main

import (
	"context"
	"time"

	"github.com/americanas-go/cache"
	"github.com/americanas-go/cache/codec/gob"
	driver "github.com/americanas-go/cache/driver/contrib/coocood/freecache.v1"
	logger "github.com/americanas-go/cache/middleware/contrib/americanas-go/log.v1"
	"github.com/americanas-go/log"
	"github.com/americanas-go/log/contrib/rs/zerolog.v1"
	"github.com/coocood/freecache"
)

func main() {

	zerolog.NewLogger(zerolog.WithLevel("TRACE"))

	fc := freecache.NewCache(1)

	drv := driver.New(fc, &driver.Options{
		TTL: 10 * time.Minute,
	})

	manager := cache.NewManager[string]("foo", gob.New[string](), drv)
	manager.Use(logger.New[string]())

	ctx := context.Background()

	if err := manager.Set(ctx, "key", "value"); err != nil {
		panic(err)
	}

	ok, value, err := manager.Get(ctx, "key")
	if err != nil {
		panic(err)
	}

	if !ok {
		log.Infof("no key found")
	}

	value2, err := manager.GetOrSet(ctx, "key2", func(ctx context.Context) (string, error) {
		log.FromContext(ctx).Infof("executed get or set")
		return "get or set", nil
	}, cache.SaveEmpty)
	if err != nil {
		panic(err)
	}

	log.Infof(value)
	log.Infof(value2)
}
