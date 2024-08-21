package main

import(
	"netarsenal/pkg/network"

	"fmt"
	"flag"
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

	for i:=0; i<=100; i++ {
		// print the current port being scanned
		fmt.Print("\b\b\b")
		fmt.Print(i)

		//network.TcpFinConnection(cliArgs.ipAdd, i)
		network.TcpConnection(cliArgs.ipAdd, i)
	}
}
