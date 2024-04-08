package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"os"
)

type EPB struct {
	BlockType            uint32 // 4
	BlockTotalLength     uint32 // 4
	InterfaceID          uint32 // 4
	TimestampHigh        uint32 // 4
	TimestampLow         uint32 // 4
	CapturedPacketLength uint32 // 4
	OriginalPacketLength uint32 // 4
	PacketData           []byte
	Options              []byte
	BTL                  uint32 // 4
}

var epbCount = 1

var protocolCount map[string]int = make(map[string]int)
var tcpSourceCount map[string]int = make(map[string]int)

func readEPB(file *os.File) (*EPB, error) {

	var endian binary.ByteOrder = binary.LittleEndian

	epb := &EPB{}

	fields := []interface{}{
		&epb.BlockType,
		&epb.BlockTotalLength,
		&epb.InterfaceID,
		&epb.TimestampHigh,
		&epb.TimestampLow,
		&epb.CapturedPacketLength,
		&epb.OriginalPacketLength,
	}

	for _, field := range fields {
		if err := binary.Read(file, endian, field); err != nil {
			return nil, err
		}
	}

	dataLen := DATALEN(int64(epb.CapturedPacketLength))

	if dataLen > 0 {
		var err error
		epb.PacketData, err = readChunk(file, dataLen)
		if err != nil {
			return nil, err
		}
	}

	optLen := int64(epb.BlockTotalLength) - 32 - dataLen

	if optLen > 0 {
		var err error
		epb.Options, err = readChunk(file, optLen)
		if err != nil {
			return nil, err
		}
	}

	if err := binary.Read(file, endian, &epb.BTL); err != nil {
		return nil, err
	}

	return epb, nil
}

func DATALEN(length int64) int64 {
	return (length + 3) & ^0x3
}

func (epb *EPB) Log() {
	log.Printf("%d - Enhanced Packet Block (EPB):", epbCount)

	// log.Printf("\tBlock Type: %x\n", epb.BlockType)
	// log.Printf("\tBlock Total Length: %d\n", epb.BlockTotalLength)
	// log.Printf("\tInterfaceID: %d\n", epb.InterfaceID)
	// log.Printf("\tTimestampHigh: %d\n", epb.TimestampHigh)
	// log.Printf("\tTimestampLow: %d\n", epb.TimestampLow)
	// log.Printf("\tCapturedPacketLength: %d\n", epb.CapturedPacketLength)
	// log.Printf("\tOriginalPacketLength: %d\n", epb.OriginalPacketLength)
	checkPacketType(epb.PacketData, epbCount)
	// log.Println("\tOptions: Not Parsed for Now")
	// log.Printf("\tBTL: %d\n", epb.BTL)
	epbCount++
}

var ethTypes = make(map[uint16]int)

func checkPacketType(packetData []byte, num int) {
	if len(packetData) < 14 {
		log.Println("\tError: Packet too short to determine type")
		return
	}

	etherType := binary.BigEndian.Uint16(packetData[12:14])
	countEthTypes(etherType)

	if etherType == 39 {
		fmt.Printf("%d, %v\n", num, packetData)
	}

	switch etherType {
	case 0x0800:
		getIp(packetData)
		getProtocol(packetData)
	case 0x0806:
		getArpDetails(packetData)
	case 0x86DD: // IPv6
		getIPv6Protocol(packetData)
	case 0x88CC: // LLDP
		logLLDPDetails(packetData)
	default:
		checkForSTP(packetData, etherType)
	}
}

func getEthernetType(eType uint16) string {
	if name, exists := EthernetTypes[eType]; exists {
		return name
	}
	log.Println(eType)

	return "Unknown"
}

func countEthTypes(eType uint16) {
	ethTypes[eType]++
}

func getIp(packetData []byte) {
	if len(packetData) >= 34 { // 14 bytes Ethernet + 20 bytes minimum IP header
		senderIP := fmt.Sprintf("%d.%d.%d.%d", packetData[26], packetData[27], packetData[28], packetData[29])
		tcpSourceCount[senderIP]++ // Count TCP packets by source IP
		//log.Printf("\tSender IP: %s\n", senderIP)
	}
}

func getProtocol(packetData []byte) {
	if len(packetData) >= 23 && binary.BigEndian.Uint16(packetData[12:14]) == 0x0800 { // Check if it's IPv4
		protocol := packetData[23] // 14 bytes Ethernet header + 9 bytes to the protocol field
		protocolName := getProtocolName(protocol)
		protocolCount[protocolName]++ // Count each protocol occurrence
		//log.Printf("\tProtocol: %s (%d)\n", protocolName, protocol)

		if protocol == 6 { // TCP protocol number
			getIp(packetData) // Call getIp specifically for TCP packets
		}
	}
}

func getProtocolName(protocol byte) string {
	if name, exists := ProtocolMap[protocol]; exists {
		return name
	}
	return "Unknown"
}

func getArpDetails(packetData []byte) {
	if len(packetData) >= 42 { // 14 bytes Ethernet header + 28 bytes ARP header
		senderMac := packetData[22:28]
		senderIP := packetData[28:32]
		targetMac := packetData[32:38]
		targetIP := packetData[38:42]
		log.Printf("\tARP - Sender MAC: %02x:%02x:%02x:%02x:%02x:%02x, Sender IP: %d.%d.%d.%d\n",
			senderMac[0], senderMac[1], senderMac[2], senderMac[3], senderMac[4], senderMac[5],
			senderIP[0], senderIP[1], senderIP[2], senderIP[3])
		log.Printf("\tARP - Target MAC: %02x:%02x:%02x:%02x:%02x:%02x, Target IP: %d.%d.%d.%d\n",
			targetMac[0], targetMac[1], targetMac[2], targetMac[3], targetMac[4], targetMac[5],
			targetIP[0], targetIP[1], targetIP[2], targetIP[3])
	}
}

func checkForSTP(packetData []byte, etherType uint16) {
	// Check destination MAC for common STP addresses
	destinationMac := fmt.Sprintf("%02x:%02x:%02x:%02x:%02x:%02x",
		packetData[0], packetData[1], packetData[2], packetData[3], packetData[4], packetData[5])

	switch destinationMac {
	case "01:80:c2:00:00:00", "01:00:0c:cc:cc:cd":
		log.Println("\tDetected STP Packet")
	default:
		log.Printf("\tUnknown Ethernet Type: %x\n", etherType)
	}
}

func getIPv6Protocol(packetData []byte) {
	if len(packetData) < 54 { // 14 bytes Ethernet header + 40 bytes IPv6 header
		log.Println("\tError: IPv6 packet too short to extract headers")
		return
	}

	nextHeader := packetData[20] // Next Header field in the IPv6 header
	if nextHeader == 58 {        // 58 is the Next Header value for ICMPv6
		logICMPv6Details(packetData)
	}
}

func logICMPv6Details(packetData []byte) {
	// 14 bytes Ethernet header + 40 bytes IPv6 header
	// ICMPv6 messages start at byte 54
	if len(packetData) < 54 {
		log.Println("\tError: Packet too short to log ICMPv6 details")
		return
	}

	icmpType := packetData[54] // ICMPv6 Type
	icmpCode := packetData[55] // ICMPv6 Code

	log.Printf("\tICMPv6 Type: %d, Code: %d\n", icmpType, icmpCode)
}

func logLLDPDetails(packetData []byte) {
	if len(packetData) < 14 {
		log.Println("\tError: LLDP packet too short to process")
		return
	}

	// LLDP packet processing could be expanded here. For now, we log a simple message.
	log.Println("\tDetected LLDP Packet")
}
