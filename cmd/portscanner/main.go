package main

import(
	"netarsenal/pkg/network"

	"net"
	"fmt"
	"flag"
	"time"
	"strconv"
	"io"
	"bytes"
)

type Flags struct {
	ipAdd string
}

func (f *Flags) commandLineFlags() {
	flag.StringVar(&f.ipAdd, "ip", "", "Provide IP address")
	flag.Parse()
}

// todo
// try to send either UDP packet or tcp fin packet for faster scan
// use goroutines for concurrently scanning ports.

func main() {
	var cliArgs Flags
	cliArgs.commandLineFlags()

	fmt.Println("Scanning ports for " + cliArgs.ipAdd)
	//udpAddr, _ := net.ResolveUDPAddr("udp", cliArgs.ipAdd)

	for i:=0; i<1000; i++ {
		//conn, err := net.Dial("tcp", cliArgs.ipAdd + ":" + strconv.Itoa(i))
		conn, err := network.scan.tcpFinConnection(cliArgs.ipAdd, strconv.Itoa(i))
		
		// print the current port being scanned
		fmt.Print("\b\b\b")
		fmt.Print(i)

		if err != nil {
			continue
		}
		defer conn.Close()
	
		// set a read deadline
		conn.SetReadDeadline(time.Now().Add(5 * time.Second))

		var buf bytes.Buffer
		io.Copy(&buf, conn)
		if buf.Len() > 0 {
			fmt.Println()
		}
		
	}
}
