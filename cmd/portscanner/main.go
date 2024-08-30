package main

import(
	"netarsenal/pkg/network"

	"fmt"
	"flag"
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

func main() {
	startTime := time.Now()

	var cliArgs Flags
	cliArgs.commandLineFlags()

	portsOpen := &network.PortsOpen {
		Data: make([]int, 0), 
	}
	var wg sync.WaitGroup

	fmt.Println("Scanning ports for " + cliArgs.ipAdd)

	for i:=0; i<=100; i++ {
		wg.Add(1)
		switch cliArgs.rType {
		case "tcp":
			go network.TcpConnection(cliArgs.ipAdd, i, portsOpen, &wg)
		case "tcpFin":
			go network.TcpFinConnection(cliArgs.ipAdd, i, portsOpen, &wg)
		default:
			fmt.Println(cliArgs.rType, "is not implemented yet.")
			return
		}
	}
	
	wg.Wait()
	portsOpen.RLock()
	fmt.Println(portsOpen.Data)
	portsOpen.RUnlock()

	fmt.Println("time taken", time.Now().Sub(startTime))
}
