package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

type Block interface {
	Log()
}

var blockCount = 0

var blockTypes map[string]int = make(map[string]int)

func countNlockTypes(bType string) {
	blockTypes[bType]++
}

func analyzeBlocks(file *os.File) {

	log.Printf("Analyzing file\n")

	for {
		blockType, err := readBlockType(file)
		if err != nil {
			break
		}
		var block Block

		switch blockType {
		case "0A0D0D0A":
			block, err = readSHB(file)

		case "01000000":
			block, err = readIDB(file)

		case "06000000":
			block, err = readEPB(file)

		case "05000000":
			block, err = readISB(file)

		default:
			log.Println("Invalid Block.")
		}

		if err != nil {
			panic(
				fmt.Sprintf("Error reading block:%s - %v\n", blockType, err),
			)
		}
		countNlockTypes(blockType)
		blockCount++

		block.Log()
	}
}

func readBlockType(file *os.File) (string, error) {

	blockType := make([]byte, 4)

	if _, err := file.Read(blockType); err != nil {
		return "", err
	}

	_, err := file.Seek(-4, io.SeekCurrent)
	if err != nil {
		panic(err)
	}

	bType := fmt.Sprintf("%X", blockType)

	//fmt.Println(bType)

	return bType, nil

}

func isSupported(blockType string) bool {

	if blockType != "0A0D0D0A" {
		log.Printf("Block:\t%s not supported!\n", blockType)
		return false
	}

	log.Printf("Magic #:\t%s\n", blockType)
	log.Println()
	return true
}
