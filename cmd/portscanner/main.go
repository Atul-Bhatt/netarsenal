package main

import(
	"netarsenal/pkg/network"

	"fmt"
	"flag"
	"sync"
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
	portsOpen := make([]int, 100)

	var wg sync.WaitGroup

	fmt.Println("Scanning ports for " + cliArgs.ipAdd)

	for i:=0; i<=100; i++ {
		wg.Add(1)	
		//network.TcpFinConnection(cliArgs.ipAdd, i)
		go network.TcpConnection(cliArgs.ipAdd, i, portsOpen, wg)
	}
	
	wg.Wait()
	fmt.Println(portsOpen)
}
