package main

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
}
