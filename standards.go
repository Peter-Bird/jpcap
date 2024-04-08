package main

var LinkTypes = map[int]string{
	0:   "NULL",                   // BSD loopback encapsulation
	1:   "ETHERNET",               // Ethernet (10Mb, 100Mb, 1Gb, 10Gb, etc)
	6:   "IEEE802_5",              // IEEE 802.5 Token Ring
	7:   "ARCNET",                 // ARCNET
	8:   "SLIP",                   // SLIP
	9:   "PPP",                    // Point-to-Point Protocol
	10:  "FDDI",                   // FDDI
	50:  "PPP_HDLC",               // PPP in HDLC-like framing
	51:  "PPP_ETHER",              // PPPoE
	100: "ATM_RFC1483",            // LLC/SNAP-encapsulated ATM
	101: "RAW",                    // Raw IP
	104: "C_HDLC",                 // Cisco HDLC
	105: "IEEE802_11",             // IEEE 802.11 wireless
	108: "FRELAY",                 // Frame Relay
	113: "LOOP",                   // Loopback
	114: "LINUX_SLL",              // Linux "cooked" capture
	117: "LTALK",                  // Apple LocalTalk
	119: "PFLOG",                  // OpenBSD pflog
	122: "IEEE802_11_PRISM",       // Prism-header 802.11
	127: "IP_OVER_FC",             // IP over Fibre Channel
	129: "SUNATM",                 // SunATM
	138: "IEEE802_11_RADIOTAP",    // Radiotap-header 802.11
	139: "ARCNET_LINUX",           // ARCNET Linux
	143: "APPLE_IP_OVER_IEEE1394", // Apple IP-over-IEEE 1394 cooked header
	144: "MTP2_WITH_PHDR",         // SS7 MTP2 with Pseudo-header
	145: "MTP2",                   // SS7 MTP2
	146: "MTP3",                   // SS7 MTP3
	147: "SCCP",                   // SS7 SCCP
	148: "DOCSIS",                 // DOCSIS
	149: "LINUX_IRDA",             // Linux-IrDA
	// more types can be added as needed
}

var EthernetTypes = map[uint16]string{
	0x0027: "JSTP",           // Juniper Spanning Tree Protocol
	0x0800: "IPv4",           // Internet Protocol version 4
	0x0806: "ARP",            // Address Resolution Protocol
	0x86DD: "IPv6",           // Internet Protocol version 6
	0x88CC: "LLDP",           // Link Layer Discovery Protocol
	0x8847: "MPLS",           // MultiProtocol Label Switching
	0x8035: "RARP",           // Reverse Address Resolution Protocol
	0x8100: "VLAN",           // Virtual Local Area Network (IEEE 802.1Q)
	0x8864: "PPPoE_S",        // PPPoE Session Stage
	0x8863: "PPPoE_D",        // PPPoE Discovery Stage
	0x88E5: "MACsec",         // IEEE 802.1AE MAC security (MACsec)
	0x0805: "X.25",           // X.25 over Ethernet
	0x814C: "SNMP",           // Simple Network Management Protocol
	0x886D: "ISO-IP",         // ISO Internet Protocol
	0x88F7: "PTP",            // Precision Time Protocol (IEEE 1588)
	0x88A8: "QinQ",           // Stacked VLANs (Q-in-Q)
	0x8808: "EFC",            // Ethernet Flow Control
	0x8809: "Slow_Protocols", // Slow Protocols such as LACP (Link Aggregation Control Protocol)
	0x8915: "RoCE",           // RDMA over Converged Ethernet
	0x8906: "FCoE",           // Fibre Channel over Ethernet
	0x8914: "FCoE_Init",      // FCoE Initialization Protocol
	0x8902: "CFM",            // IEEE 802.1ag Connectivity Fault Management
	0x891D: "TTE",            // TTEthernet Protocol Control Frame
	0x0842: "Wake_on_LAN",    // Wake on LAN
	0x8102: "STP",            // Spanning Tree Protocol
	// more types can be added as needed
}

var ProtocolMap = map[byte]string{
	1:   "ICMP",
	2:   "IGMP",
	6:   "TCP",
	17:  "UDP",
	41:  "IPv6 Encapsulation",
	47:  "GRE",
	50:  "ESP",
	51:  "AH",
	89:  "OSPF",
	115: "L2TP",
	// more types can be added as needed
}

var IPv4Protocols = map[uint8]string{
	0:   "HOPOPT",     // IPv6 Hop-by-Hop Option
	1:   "ICMP",       // Internet Control Message Protocol
	2:   "IGMP",       // Internet Group Management Protocol
	3:   "GGP",        // Gateway-to-Gateway Protocol
	4:   "IP-in-IP",   // IP in IP (encapsulation)
	6:   "TCP",        // Transmission Control Protocol
	8:   "EGP",        // Exterior Gateway Protocol
	9:   "IGP",        // Any private interior gateway (used by Cisco for their IGRP)
	17:  "UDP",        // User Datagram Protocol
	18:  "MUX",        // Multiplexing
	27:  "RDP",        // Reliable Datagram Protocol
	29:  "ISO-TP4",    // ISO Transport Protocol Class 4
	41:  "IPv6",       // IPv6 encapsulation
	43:  "IPv6-Route", // Routing Header for IPv6
	44:  "IPv6-Frag",  // Fragment Header for IPv6
	47:  "GRE",        // Generic Routing Encapsulation
	50:  "ESP",        // Encapsulating Security Payload
	51:  "AH",         // Authentication Header
	57:  "SKIP",       // SKIP
	58:  "IPv6-ICMP",  // ICMP for IPv6
	59:  "IPv6-NoNxt", // No Next Header for IPv6
	60:  "IPv6-Opts",  // Destination Options for IPv6
	88:  "EIGRP",      // Enhanced Interior Gateway Routing Protocol
	89:  "OSPF",       // Open Shortest Path First
	94:  "IPIP",       // IP-within-IP Encapsulation Protocol
	95:  "MICP",       // Mobile Internetworking Control Protocol (not used)
	97:  "ETHERIP",    // Ethernet-within-IP Encapsulation
	98:  "ENCAP",      // Encapsulation Header
	103: "PIM",        // Protocol Independent Multicast
	112: "VRRP",       // Virtual Router Redundancy Protocol
	115: "L2TP",       // Layer Two Tunneling Protocol
	132: "SCTP",       // Stream Control Transmission Protocol
	133: "FC",         // Fibre Channel
	137: "MPLS-in-IP", // MPLS-in-IP
	139: "HIP",        // Host Identity Protocol
	140: "SHIM6",      // Shim6 Protocol
	142: "WESP",       // Wrapped Encapsulating Security Payload
	143: "ROHC",       // Robust Header Compression
	// Additional protocol numbers can be added as necessary
}

var UDPProtocols = map[int]string{
	53:   "DNS",             // Domain Name System
	67:   "DHCP_Server",     // Dynamic Host Configuration Protocol (Server)
	68:   "DHCP_Client",     // Dynamic Host Configuration Protocol (Client)
	69:   "TFTP",            // Trivial File Transfer Protocol
	123:  "NTP",             // Network Time Protocol
	161:  "SNMP",            // Simple Network Management Protocol
	162:  "SNMP_Trap",       // SNMP Trap
	500:  "IKE",             // Internet Key Exchange for IPsec
	514:  "Syslog",          // Syslog used for system logging
	520:  "RIP",             // Routing Information Protocol
	137:  "NetBIOS_NS",      // NetBIOS Name Service
	138:  "NetBIOS_DGM",     // NetBIOS Datagram Service
	139:  "NetBIOS_SSN",     // NetBIOS Session Service
	445:  "SMB",             // Microsoft SMB Protocol
	1194: "OpenVPN",         // OpenVPN
	1701: "L2TP",            // Layer 2 Tunneling Protocol
	1812: "RADIUS",          // RADIUS authentication protocol
	1813: "RADIUS_Acct",     // RADIUS accounting
	2049: "NFS",             // Network File System
	3478: "STUN",            // Session Traversal Utilities for NAT
	5060: "SIP",             // Session Initiation Protocol (VoIP signaling)
	5061: "SIPS",            // Secure Session Initiation Protocol
	5600: "Radmin",          // Radmin remote administration tool
	6881: "BitTorrent",      // BitTorrent
	3544: "Teredo",          // Teredo tunneling
	3389: "RDP",             // Remote Desktop Protocol
	1900: "SSDP",            // Simple Service Discovery Protocol
	4500: "IPSec_NAT-T",     // IPsec NAT Traversal
	119:  "NNTP",            // Network News Transfer Protocol
	1434: "MS-SQL",          // Microsoft SQL Server
	1645: "RADIUS_alt",      // Alternate RADIUS Authentication (historical)
	1646: "RADIUS_Acct_alt", // Alternate RADIUS Accounting (historical)
	5353: "mDNS",            // Multicast DNS
	4789: "VXLAN",           // Virtual Extensible LAN
	989:  "FTPS_data",       // FTP over TLS/SSL (data)
	990:  "FTPS",            // FTP over TLS/SSL (control)
	// Additional UDP protocols and ports can be added as necessary
}
