package network

import(
	"net"
	"log"
	
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

func getIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

func tcpFinConnection(ip_addr, port string) {
	// Open a handle on the network interface to send the packet
	handle, err := pcap.OpenLive("wlan0", 65535, true, pcap.BlockForever)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	// create custom packet with FIN bit set
	ipLayer := &layers.IPv4{
		SrcIP: getIP(), 
		DstIP: net.IP(ip_addr), 
		Protocol: layers.IPProtocolTCP,
	}

	tcpLayer := &layers.TCP{
		SrcPort: layers.TCPPort(12345),
		DstPort: layers.TCPPort(port),
		SYN: false,
		FIN: true,
		Seq: 1234567,
	}

	tcpLayer.SetNetworkLayerForChecksum(ipLayer)

	// Create a packet with the IP and TCP layers
	buffer := gopacket.NewSerializeBuffer()
	options := gopacket.SerializeOptions{
		ComputeChecksums: true,
		FixLengths: true,
	}
	err = gopacket.SerializeLayers(buffer, options, ipLayer, tcpLayer)
	if err != nil {
		log.Fatal(err)
	}

	// Send the packet
	err = handle.WriteaPacketData(buffer.Bytes())
	
}


