package main

import (
	"context"
	"fmt"
	"time"

	"github.com/americanas-go/cache"
	"github.com/americanas-go/cache/codec/gob"
	driver "github.com/americanas-go/cache/driver/contrib/coocood/freecache.v1"
	"github.com/coocood/freecache"
)

func main() {

	fc := freecache.NewCache(1)

	drv := driver.New(fc, &driver.Options{
		TTL: 10 * time.Minute,
	})

	manager := cache.NewManager[string]("foo", gob.New[string](), drv)

	ctx := context.Background()

	if err := manager.Set(ctx, "key", "value"); err != nil {
		panic(err)
	}

	ok, value, err := manager.Get(ctx, "key")
	if err != nil {
		panic(err)
	}

	if !ok {
		fmt.Println("no key found")
	}

	fmt.Println(value)
}
