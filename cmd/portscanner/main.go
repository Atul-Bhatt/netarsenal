package main

import(
	"net"
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

func main() {
	var cliArgs Flags
	cliArgs.commandLineFlags()

	conn, err := net.Dial("tcp", cliArgs.ipAdd)
	if err != nil {
		fmt.Println("Failed to create a connection")
		return
	}

	var b []byte
	_, readErr := conn.Read(b)
	if readErr != nil {
		fmt.Println("Failed to read from connection")
		return
	}

	fmt.Println(b)
}
