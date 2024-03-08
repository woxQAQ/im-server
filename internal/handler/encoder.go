package handler

import (
	"bytes"
	"encoding/gob"
	"sync"
)

type Encoder interface {
	Encode(data any) ([]byte, error)
	Decode(encodeData []byte, decodeData any) error
}

type gobEncoder struct {
	sync.Pool
}

func newGobEncoder() *gobEncoder {
	return &gobEncoder{
		Pool: sync.Pool{
			New: func() any {
				return new(bytes.Buffer)
			},
		},
	}
}

func (e *gobEncoder) Encode(data any) ([]byte, error) {
	var buf = e.Pool.Get().(bytes.Buffer)
	defer e.Pool.Put(buf)
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(data); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (e *gobEncoder) Decode(encodeData []byte, decodeData any) error {
	var buf = e.Pool.Get().(bytes.Buffer)
	defer e.Pool.Put(buf)
	buf.Write(encodeData)
	dec := gob.NewDecoder(&buf)
	if err := dec.Decode(decodeData); err != nil {
		return err
	}
	return nil
}
