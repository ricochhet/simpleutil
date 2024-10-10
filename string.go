package simpleutil

import (
	"bytes"
	"unicode/utf16"
	"unsafe"
)

func StringToBytes(input string) []byte {
	tmp := []byte(input)
	tmp = append(tmp, bytes.Repeat([]byte{0}, 2-(len(tmp)%2))...) //nolint:mnd // wontfix

	return tmp
}

func GetStringFromBytes(data []byte, start, end int) string {
	var contentID string

	if end > len(data) {
		end = len(data)
	}

	rawSlice := data[start:end]
	u16Slice := ((*[1 << 30]uint16)(unsafe.Pointer(&rawSlice[0])))[:len(rawSlice)/2]

	nullIndex := -1

	for i, c := range u16Slice {
		if c == 0 {
			nullIndex = i
			break
		}
	}

	if nullIndex != -1 {
		contentID = string(utf16.Decode(u16Slice[:nullIndex]))
	} else {
		contentID = string(utf16.Decode(u16Slice))
	}

	return contentID
}
