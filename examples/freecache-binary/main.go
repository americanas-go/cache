package main

import (
	"context"
	"fmt"
	"time"

	"github.com/americanas-go/cache"
	"github.com/americanas-go/cache/codec/binary"
	driver "github.com/americanas-go/cache/driver/contrib/coocood/freecache.v1"
	"github.com/coocood/freecache"
)

type packet struct {
	Sensid uint32
	Locid  uint16
	Tstamp uint32
	Temp   int16
}

func main() {

	fc := freecache.NewCache(1)

	drv := driver.New(fc, &driver.Options{
		TTL: 10 * time.Minute,
	})

	codec := binary.New[packet]()

	manager := cache.NewManager[packet]("foo", codec, drv)

	ctx := context.Background()

	data := packet{Sensid: 1, Locid: 1233, Tstamp: 123452123, Temp: 12}

	if err := manager.Set(ctx, "key", data); err != nil {
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
