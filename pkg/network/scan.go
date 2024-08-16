package network

import(
	"net"
	"log"
	
//	"github.com/google/gopacket"
//	"github.com/google/gopacket/layers"
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

func udpConnection(ip_addr, port string) {
	// create packet
//	ip := &layers.IPv4{
//		SrcIP: 
//	}
}


