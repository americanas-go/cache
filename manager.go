package cache

import (
	"bytes"
	"context"
	"encoding/gob"
)

type Manager[T any] struct {
	driver Driver
}

func (m *Manager[T]) Get(ctx context.Context, key string) (ok bool, data T, err error) {

	var b []byte
	b, err = m.driver.Get(ctx, key)
	if err != nil {
		if err.Error() == "Entry not found" {
			return false, data, nil
		}
		return false, data, err
	}

	if len(b) > 0 {
		var buf bytes.Buffer
		dec := gob.NewDecoder(&buf)
		buf.Write(b)
		if err = dec.Decode(&data); err != nil {
			return false, data, err
		}
		return true, data, err
	}

	return false, data, err
}

func (m *Manager[T]) Save(ctx context.Context, key string, data T, opts ...OptionSet) (err error) {

	opt := Option{
		SaveEmpty: false,
		AsyncSave: false,
	}

	for _, o := range opts {
		o()(&opt)
	}

	var buf bytes.Buffer

	enc := gob.NewEncoder(&buf)
	if err = enc.Encode(data); err != nil {
		return err
	}

	b := buf.Bytes()

	if len(b) > 0 || opt.SaveEmpty {

		if opt.AsyncSave {

			go func(ctx context.Context, key string, b []byte) {
				m.driver.Set(ctx, key, b)
			}(ctx, key, b)

		} else {
			err = m.driver.Set(ctx, key, b)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (m *Manager[T]) GetOrSave(ctx context.Context, key string, cacheable Cacheable[T], opts ...OptionSet) (data T, err error) {

	var ok bool
	ok, data, err = m.Get(ctx, key)
	if err != nil || ok {
		return data, err
	}

	if !ok {

		data, err = cacheable(ctx)
		if err != nil {
			return data, err
		}

		err = m.Save(ctx, key, data, opts...)
		if err != nil {
			return data, err
		}

	}

	return data, err
}

func NewManager[T any](driver Driver) *Manager[T] {
	return &Manager[T]{driver: driver}
}
