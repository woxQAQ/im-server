package convert

import (
	"bytes"
	"encoding/binary"
)

func BytesToInt64(data []byte) (int64, error) {
	var res int64
	err := binary.Read(bytes.NewReader(data), binary.LittleEndian, &res)
	return res, err
}
