package gateway

import (
	"bytes"
	"encoding/gob"
)

type Encoder interface {
	Encode(data any) ([]byte, error)
	Decode(encodeData []byte, decodeData any) error
}

type gobEncoder struct{}

func newGobEncoder() *gobEncoder {
	return &gobEncoder{}
}

func (e *gobEncoder) Encode(data any) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(data); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (e *gobEncoder) Decode(encodeData []byte, decodeData any) error {
	var buf bytes.Buffer
	buf.Write(encodeData)
	dec := gob.NewDecoder(&buf)
	if err := dec.Decode(decodeData); err != nil {
		return err
	}
	return nil
}
