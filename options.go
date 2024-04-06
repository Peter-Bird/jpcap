package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

type Option struct {
	Code   uint16
	Length uint16
	Value  []byte
}

// readOptions reads 'length' bytes from the provided file.
func readChunk(file *os.File, length int64) ([]byte, error) {
	options := make([]byte, length)
	if _, err := io.ReadFull(file, options); err != nil {
		return nil, err
	}
	return options, nil
}

func ParseOptions(data []byte) ([]Option, error) {
	var options []Option

	currentIndex := 0

	for currentIndex < len(data) {

		if len(data)-currentIndex < 4 {
			return nil, fmt.Errorf("incomplete option at index %d", currentIndex)
		}

		code := binary.LittleEndian.Uint16(data[currentIndex : currentIndex+2])
		length := binary.LittleEndian.Uint16(data[currentIndex+2 : currentIndex+4])

		if code == 0 && length == 0 {
			break // End of options
		}

		if int(length) > len(data)-(currentIndex+4) {
			return nil, fmt.Errorf("option at index %d has length beyond data bounds", currentIndex)
		}

		value := data[currentIndex+4 : currentIndex+4+int(length)]

		options = append(options, Option{Code: code, Length: length, Value: value})

		currentIndex += 4 + int(length)

		// Options are padded to 4-byte boundaries
		if padding := currentIndex % 4; padding != 0 {
			currentIndex += 4 - padding
		}
	}

	return options, nil
}
