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

func main() {
	var cliArgs Flags
	cliArgs.commandLineFlags()
	portsOpen := &network.PortsOpen {
		Data: make([]int, 1000), 
	}
	var wg sync.WaitGroup

	fmt.Println("Scanning ports for " + cliArgs.ipAdd)

	for i:=0; i<=60; i++ {
		wg.Add(1)
		go network.TcpConnection(cliArgs.ipAdd, i, portsOpen, &wg)
	}
	
	wg.Wait()
	portsOpen.RLock()
	fmt.Println(portsOpen)
	portsOpen.RLock()
}
