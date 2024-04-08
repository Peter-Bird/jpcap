package main

import (
	"fmt"
	"log"
	"os"
)

// Keep
const (
	Application = "jpcap"
	Version     = "1.0.0"
)

// Keep
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

// Keep
func usage() error {
	log.Printf("Executing:\t%s\n", Application)
	log.Printf("Version:\t%s\n", Version)

	if len(os.Args) != 2 {
		return fmt.Errorf("Usage: go run . <pcapng_file>")
	}
	return nil
}

// Keep
func getFile() (*os.File, error) {
	fileName := os.Args[1]
	log.Printf("Opening:\t%s\n", fileName)
	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	return file, nil
}

// Keep
func printStatistics() {
	log.Println()

	log.Printf("Total Blocks: %17d\n", blockCount)
	log.Println()

	log.Println("Block Types:")
	for blockType, count := range blockTypes {
		log.Printf("\t%-15s: %10d\n", blockType, count)
	}
	log.Println()

	log.Println("Ethernet Types:")
	for ethType, count := range ethTypes {
		log.Printf("\t%-15s: %10d\n", getEthernetType(ethType), count)
	}
	log.Println()

	log.Println("IPV4 Counts:")
	for protocol, count := range protocolCount {
		log.Printf("\t%-15s: %10d\n", protocol, count)
	}
	log.Println()

	log.Println("UDP Counts:")
	for protocol, count := range udpTypes {
		log.Printf("\t%-15s: %10d\n", protocol, count)
	}
	log.Println()

	log.Println("TCP/IP Packets by Source IP:")
	for ip, count := range tcpSourceCount {
		log.Printf("\t%-15s: %10d\n", ip, count)
	}
	log.Println()
}
