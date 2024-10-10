package simpleutil

import (
	"bytes"
	"fmt"
)

func HexStringToBytes(hexStr string) ([]byte, error) {
	var bytes []byte

	for i := 0; i < len(hexStr); i += 2 {
		var newByte byte

		_, err := fmt.Sscanf(hexStr[i:i+2], "%02X", &newByte)
		if err != nil {
			return nil, err
		}

		bytes = append(bytes, newByte)
	}

	return bytes, nil
}

func FindAllByteOccurrences(data []byte, pattern []byte) []int {
	var indices []int

	for i := range data {
		if bytes.HasPrefix(data[i:], pattern) {
			indices = append(indices, i)
		}
	}

	return indices
}

func ReplaceByteOccurrences(original []byte, expectedBytes []byte, replacement []byte, occurrenceToReplace int) []byte {
	var result []byte

	remaining := original
	occurrenceCount := 0

	for {
		index := bytes.Index(remaining, expectedBytes)
		if index == -1 {
			result = append(result, remaining...)
			break
		}

		result = append(result, remaining[:index]...)

		occurrenceCount++

		if occurrenceToReplace == 0 || occurrenceCount == occurrenceToReplace {
			replacementLen := len(replacement)

			if replacementLen > len(expectedBytes) {
				replacementLen = len(expectedBytes)
			}

			result = append(result, replacement[:replacementLen]...)
		} else {
			result = append(result, remaining[index:index+len(expectedBytes)]...)
		}

		remaining = remaining[index+len(expectedBytes):]
	}

	return result
}
