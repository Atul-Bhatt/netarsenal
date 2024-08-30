package network

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"sync"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

type PortsOpen struct {
	sync.RWMutex
	Data []int
}

func getIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

func TcpFinConnection(ip_addr string, port int, _portsOpen *PortsOpen, wg *sync.WaitGroup) {
	// Open a handle on the network interface to send the packet
	handle, err := pcap.OpenLive(`\Device\NPF_{43C5B23F-AE4F-46FD-92BB-AD3A15667174}`, 65535, true, pcap.BlockForever)
	if err != nil {
		log.Fatal("Error in pcap.OpenLive: ", err)
	}
	defer wg.Done()
	defer handle.Close()

	// create custom packet with FIN bit set
	ipLayer := &layers.IPv4{
		SrcIP:    getIP(),
		DstIP:    net.ParseIP(ip_addr),
		Protocol: layers.IPProtocolTCP,
	}

	tcpLayer := &layers.TCP{
		SrcPort: layers.TCPPort(12345),
		DstPort: layers.TCPPort(port),
		SYN:     false,
		FIN:     true,
		Seq:     1234567,
	}

	tcpLayer.SetNetworkLayerForChecksum(ipLayer)

	// Create a packet with the IP and TCP layers
	buffer := gopacket.NewSerializeBuffer()
	options := gopacket.SerializeOptions{
		ComputeChecksums: true,
		FixLengths:       true,
	}
	err = gopacket.SerializeLayers(buffer, options, ipLayer, tcpLayer)
	if err != nil {
		log.Fatal("Error in gopacket.SerializeLayers: ", err)
	}

	// Send the packet
	err = handle.WritePacketData(buffer.Bytes())

	if err != nil {
		log.Fatal("Error while writing packet: ", err)
	}

	readBuff, _, readErr := handle.ReadPacketData()
	if readErr != nil {
		log.Fatal("Error while reading packet: ", readErr)
	}

	fmt.Println(readBuff)
}

func TcpConnection(ip_addr string, port int, portsOpen *PortsOpen, wg *sync.WaitGroup) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", ip_addr, port))

	defer wg.Done()

	if err != nil {
		return
	}

	// set connection timeout
	conn.SetReadDeadline(time.Now().Add(100 * time.Millisecond))

	var buf bytes.Buffer
	io.Copy(&buf, conn)

	//if buf.Len() > 0 {
	portsOpen.Lock()
	portsOpen.Data = append(portsOpen.Data, port)
	portsOpen.Unlock()
	//}
}
