package main

import (
	"fmt"
	"log"
	"os"
)

const (
	Application = "jpcap"
	Version     = "1.0.0"
)

func main() {

	if err := usage(); err != nil {
		log.Fatal(err)
	}

	file, err := getFile()
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	magic, _ := readBlockType(file)
	if isSupported(magic) {
		analyzeBlocks(file)
	} else {
		log.Printf("Exiting.\n")
	}

	printStatistics() // Call after processing all packets
}

func usage() error {
	log.Printf("Executing:\t%s\n", Application)
	log.Printf("Version:\t%s\n", Version)

	if len(os.Args) != 2 {
		return fmt.Errorf("Usage: go run . <pcapng_file>")
	}
	return nil
}

func getFile() (*os.File, error) {
	fileName := os.Args[1]
	log.Printf("Opening:\t%s\n", fileName)
	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	return file, nil
}

func printStatistics() {
	log.Println("Protocol Counts:")
	for protocol, count := range protocolCount {
		log.Printf("\t%s: %d\n", protocol, count)
	}

	log.Println("TCP Packets by Source IP:")
	for ip, count := range tcpSourceCount {
		log.Printf("\t%s: %d\n", ip, count)
	}
}
