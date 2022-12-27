package main

import (
	"context"
	"time"

	"github.com/americanas-go/cache"
	"github.com/americanas-go/cache/codec/gob"
	driver "github.com/americanas-go/cache/driver/contrib/coocood/freecache.v1"
	"github.com/americanas-go/log"
	"github.com/americanas-go/log/contrib/rs/zerolog.v1"
	"github.com/coocood/freecache"
)

func main() {

	ctx := context.Background()

	zerolog.NewLogger(zerolog.WithLevel("TRACE"))

	codec := gob.New[string]()

	fc1 := freecache.NewCache(1)
	fc2 := freecache.NewCache(1)

	drv1 := driver.New(fc1, &driver.Options{
		TTL: 10 * time.Minute,
	})

	drv2 := driver.New(fc2, &driver.Options{
		TTL: 10 * time.Minute,
	})

	v, _ := codec.Encode("value drv2")

	if err := drv2.Set(ctx, "key", v); err != nil {
		panic(err)
	}

	manager := cache.NewManager[string]("foo", codec, drv1, drv2)

	value, err := manager.GetOrSet(ctx, "key", func(ctx context.Context) (string, error) {
		return "value", nil
	})
	if err != nil {
		panic(err)
	}

	log.Infof("returned: %s", value)

	log.Infof("entries on drv1: %v", fc1.EntryCount())
	log.Infof("entries on drv2: %v", fc2.EntryCount())
}
