package gob

import (
	"bytes"
	"encoding/binary"
)

type Codec[T any] struct {
}

func (c *Codec[T]) Encode(e T) (b []byte, err error) {
	var buf bytes.Buffer

	if err = binary.Write(&buf, binary.LittleEndian, e); err != nil {
		return b, err
	}

	b = buf.Bytes()
	return b, err
}

func (c *Codec[T]) Decode(b []byte, data *T) (err error) {
	buf := bytes.NewReader(b)
	return binary.Read(buf, binary.LittleEndian, &data)
}

func New[T any]() *Codec[T] {
	return &Codec[T]{}
}
