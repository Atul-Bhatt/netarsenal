package main

import(
	"net"
	"fmt"
	"flag"
	"time"
)

const PORT = "80"

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

	conn, err := net.Dial("tcp", cliArgs.ipAdd + ":" +  PORT)
	if err != nil {
		fmt.Println("Failed to create a connection")
		return
	}
	defer conn.Close()

	fmt.Println("Scanning ports for " + cliArgs.ipAdd)

	// set a read deadline
	conn.SetReadDeadline(time.Now().Add(5 * time.Second))

	b := make([]byte, 1024)
	readFromConn, readErr := conn.Read(b)
	if readErr != nil {
		fmt.Println("Failed to read from connection")
		return
	}

	fmt.Println(b, readFromConn)
}
