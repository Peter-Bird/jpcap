package main

import (
	"encoding/binary"
	"os"
)

// Interface Description Block
type IDB struct {
	BlockType        uint32
	BlockTotalLength uint32
	LinkType         uint16
	Reserved         uint16 // Must be zero
	SnapLen          uint32
	Options          []byte
	BTL              uint32
}

func readIDB(file *os.File) (*IDB, error) {

	var endian binary.ByteOrder = binary.LittleEndian

	idb := &IDB{}

	fields := []interface{}{
		&idb.BlockType,
		&idb.BlockTotalLength,
		&idb.LinkType,
		&idb.Reserved,
		&idb.SnapLen,
	}

	for _, field := range fields {
		if err := binary.Read(file, endian, field); err != nil {
			return nil, err
		}
	}

	optLen := int64(idb.BlockTotalLength) - 20

	if optLen > 0 {
		idb.Options = make([]byte, optLen)
		if _, err := file.Read(idb.Options); err != nil {
			return nil, err
		}
	}
	if err := binary.Read(file, endian, &idb.BTL); err != nil {
		return nil, err
	}

	return idb, nil
}

func (idb *IDB) Log() {
	// log.Println("Interface Description Block (IDB):")
	// log.Printf("\tBlock Type: %x\n", idb.BlockType)
	// log.Printf("\tBlock Total Length: %d\n", idb.BlockTotalLength)
	// log.Printf("\tLinkType: %x\n", idb.LinkType)
	// log.Printf("\tReserved: %d\n", idb.Reserved)
	// log.Printf("\tSnapLen: %d\n", idb.SnapLen)
	// log.Println("\tOptions: Not Parsed for Now")
	// log.Printf("\tBTL: %d\n", idb.BTL)
}
