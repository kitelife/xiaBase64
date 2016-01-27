package xiaBase64

import (
	"encoding/binary"
	"fmt"
)

func Decode(input []byte) ([]byte, error) {
	byteCount := len(input)
	if float32(byteCount/4) < float32(byteCount)/4 {
		return nil, fmt.Errorf("Not valid input")
	}
	if input[byteCount-1] == '=' {
		input[byteCount-1] = 'A'
	}
	if input[byteCount-2] == '=' {
		input[byteCount-2] = 'A'
	}

	var indexes []uint32

	for _, b := range input {
		indexes = append(indexes, decodingTable[b])
	}

	beforeBase64 := []byte{}
	pos := 0
	for pos < (byteCount - 1) {
		v := indexes[pos]<<18 + indexes[pos+1]<<12 + indexes[pos+2]<<6 + indexes[pos+3]

		first := v >> 16
		second := (v << 16) >> 24
		third := (v << 24) >> 24

		buf := make([]byte, 4)
		binary.BigEndian.PutUint32(buf, first)
		beforeBase64 = append(beforeBase64, buf[3])
		binary.BigEndian.PutUint32(buf, second)
		beforeBase64 = append(beforeBase64, buf[3])
		binary.BigEndian.PutUint32(buf, third)
		beforeBase64 = append(beforeBase64, buf[3])

		pos = pos + 4
	}

	return beforeBase64, nil
}
