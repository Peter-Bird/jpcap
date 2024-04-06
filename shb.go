package main

import (
	"encoding/binary"

	//	"io"
	"log"
	"os"
)

// Section Header Block
type SHB struct {
	BlockType        uint32 // 4
	BlockTotalLength uint32 // 4
	ByteOrderMagic   uint32 // 4
	MajorVersion     uint16 // 2
	MinorVersion     uint16 // 2
	SectionLength    int64  // 8  Can be -1 for unspecified length
	Options          []byte // N  BlockTotalLength - 28
	BTL              uint32 // 4
}

func readSHB(file *os.File) (*SHB, error) {
	var endian binary.ByteOrder = binary.LittleEndian

	shb := &SHB{}

	fields := []interface{}{
		&shb.BlockType,
		&shb.BlockTotalLength,
		&shb.ByteOrderMagic, // For now assuming LittleEndian
		&shb.MajorVersion,
		&shb.MinorVersion,
		&shb.SectionLength,
	}

	for _, field := range fields {
		if err := binary.Read(file, endian, field); err != nil {
			return nil, err
		}
	}

	optLen := int64(shb.BlockTotalLength) - 28

	if optLen > 0 {
		var err error
		shb.Options, err = readChunk(file, optLen)
		if err != nil {
			return nil, err
		}
	}

	if err := binary.Read(file, endian, &shb.BTL); err != nil {
		return nil, err
	}

	return shb, nil
}

func (shb *SHB) Log() {
	log.Println("Section Header Block (SHB):")
	log.Printf("\tBlock Type: %x\n", shb.BlockType)
	log.Printf("\tBlock Total Length: %d\n", shb.BlockTotalLength)
	log.Printf("\tByte Order Magic: %x\n", shb.ByteOrderMagic)
	log.Printf("\tMajor Version: %d\n", shb.MajorVersion)
	log.Printf("\tMinor Version: %d\n", shb.MinorVersion)
	log.Printf("\tSection Length: %d\n", shb.SectionLength)
	log.Println("\tOptions: Not Parsed for Now")
	log.Printf("\tBTL: %d\n", shb.BTL)
}
