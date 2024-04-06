package main

import (
	"encoding/binary"
	"log"
	"os"
)

// Interface Statistics Block
type ISB struct {
	BlockType        uint32
	BlockTotalLength uint32
	InterfaceID      uint32
	TimestampHigh    uint32
	TimestampLow     uint32
	Options          []byte
	BTL              uint32
}

func readISB(file *os.File) (*ISB, error) {

	var endian binary.ByteOrder = binary.LittleEndian

	isb := &ISB{}

	fields := []interface{}{
		&isb.BlockType,
		&isb.BlockTotalLength,
		&isb.InterfaceID,
		&isb.TimestampHigh,
		&isb.TimestampLow,
	}

	for _, field := range fields {
		if err := binary.Read(file, endian, field); err != nil {
			return nil, err
		}
	}

	optLen := int64(isb.BlockTotalLength) - 24

	if optLen > 0 {
		isb.Options = make([]byte, optLen)
		if _, err := file.Read(isb.Options); err != nil {
			return nil, err
		}
	}
	if err := binary.Read(file, endian, &isb.BTL); err != nil {
		return nil, err
	}

	return isb, nil
}

func (isb *ISB) Log() {
	log.Println("Interface Statistics Block (ISB):")
	log.Printf("\tBlock Type: %x\n", isb.BlockType)
	log.Printf("\tBlock Total Length: %d\n", isb.BlockTotalLength)
	log.Printf("\tInterfaceID: %x\n", isb.InterfaceID)
	log.Printf("\tTimestampHigh: %d\n", isb.TimestampHigh)
	log.Printf("\tTimestampLow: %d\n", isb.TimestampLow)
	log.Println("\tOptions: Not Parsed for Now")
	log.Printf("\tBTL: %d\n", isb.BTL)
}
