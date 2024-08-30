package main

import (
	"netarsenal/pkg/network"

	"flag"
	"fmt"
	"sync"
	"time"
)

type Flags struct {
	ipAdd string
	rType string
}

func (f *Flags) commandLineFlags() {
	flag.StringVar(&f.ipAdd, "ip", "", "Provide IP address")
	flag.StringVar(&f.rType, "rt", "tcp", "Request Type (tcp, tcpSyn, tcpFin, udp, xmas, null)")
	flag.Parse()
}

const MAX_PORTS = 65535

func main() {
	startTime := time.Now()

	var cliArgs Flags
	cliArgs.commandLineFlags()

	portsOpen := &network.PortsOpen{
		Data: make([]int, 0),
	}
	var wg sync.WaitGroup

	// get the type of request from command line
	var connFunc func(string, int, *network.PortsOpen, *sync.WaitGroup)

	switch cliArgs.rType {
	case "tcp":
		connFunc = network.TcpConnection
	case "tcpFin":
		connFunc = network.TcpFinConnection
	default:
		fmt.Println(cliArgs.rType, "is not implemented yet.")
		return
	}

	fmt.Println("Scanning ports for " + cliArgs.ipAdd)
	for i := 0; i <= MAX_PORTS; i++ {
		wg.Add(1)
		go connFunc(cliArgs.ipAdd, i, portsOpen, &wg)
	}

	wg.Wait()
	portsOpen.RLock()
	fmt.Println(portsOpen.Data)
	portsOpen.RUnlock()

	fmt.Println("time taken", time.Now().Sub(startTime))
}
