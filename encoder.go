package xiaBase64

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func Encode(input []byte) []byte {
	byteCount := len(input)

	intResult := byteCount * 8 / 6
	floatResult := float32(byteCount*8) / 6

	extraNum := 0
	if float32(intResult) < floatResult {
		input = append(input, 0x0)
		extraNum = 1
		if (floatResult - float32(intResult)) < 0.5 {
			input = append(input, 0x0)
			extraNum = 2
		}
	}

	pos := 0
	byteCount = len(input)

	afterBase64 := []byte{}
	for pos < (byteCount - 1) {
		bytesReader := bytes.NewReader(append([]byte{0x0}, input[pos:pos+3]...))

		var bi uint32
		err := binary.Read(bytesReader, binary.BigEndian, &bi)
		if err != nil {
			fmt.Println("binary.Read failed:", err)
		}

		first := bi >> 18
		second := (bi << 14) >> 26
		third := (bi << 20) >> 26
		fourth := (bi << 26) >> 26

		afterBase64 = append(afterBase64,
			encodingTable[first], encodingTable[second], encodingTable[third], encodingTable[fourth])
		pos = pos + 3
	}
	lenAfterBase64 := len(afterBase64)
	if extraNum > 0 {
		afterBase64[lenAfterBase64-1] = '='
	}
	if extraNum > 1 {
		afterBase64[lenAfterBase64-2] = '='
	}
	return afterBase64
}
