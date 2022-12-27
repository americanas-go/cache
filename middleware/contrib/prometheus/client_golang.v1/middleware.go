package prometheus

import (
	"github.com/americanas-go/cache"
	"github.com/americanas-go/log"
)

type middleware[R any] struct{}

func (l middleware[R]) Del(c *cache.Context[R], s string) error {
	middleware := log.FromContext(c.GetContext())
	middleware.Tracef("executing Del method")
	defer middleware.Debugf("executed Del method")
	return c.Del(s)
}

func (l middleware[R]) Get(c *cache.Context[R], s string) ([]byte, error) {
	middleware := log.FromContext(c.GetContext())
	middleware.Tracef("executing Get method")
	defer middleware.Debugf("executed Get method")
	return c.Get(s)
}

func (l middleware[R]) Set(c *cache.Context[R], s string, data []byte) error {
	middleware := log.FromContext(c.GetContext())
	middleware.Tracef("executing Set method")
	defer middleware.Debugf("executed Set method")
	return c.Set(s, data)
}

func New[R any]() cache.Middleware[R] {
	return &middleware[R]{}
}
