package network 

import(
	"net"
	"log"
	"fmt"
	
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

func TcpFinConnection(ip_addr string, port int) {
	// Open a handle on the network interface to send the packet
	handle, err := pcap.OpenLive(`\Device\NPF_{6C1EA2F0-95C6-4859-8D8F-2EE3B1D2F3A7}`, 65535, true, pcap.BlockForever)
	if err != nil {
		log.Fatal("Error in pcap.OpenLive: ", err)
	}
	defer handle.Close()

	// create custom packet with FIN bit set
	ipLayer := &layers.IPv4{
		SrcIP: getIP(), 
		DstIP: net.ParseIP(ip_addr), 
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
		log.Fatal("Error in gopacket.SerializeLayers: ", err)
	}

	// Send the packet
	err = handle.WritePacketData(buffer.Bytes())
	if len(buffer.Bytes()) > 0 {
		fmt.Println()
	}
}


