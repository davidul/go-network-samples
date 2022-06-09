package main

import (
	"flag"
	"fmt"
	"net"
)

func main() {

	var network string
	var addr string
	flag.StringVar(&network, "n", "tcp", "network type [tcp/unix]")
	flag.StringVar(&addr, "e", ":4040", "service endpoint [addr or path]")

	flag.Parse()
	arg := flag.Arg(0)

	dial, err := net.Dial(network, addr)
	if err != nil {
		panic(err)
	}

	tmp := []byte(arg)
	write, err := dial.Write(tmp)
	if err != nil {
		fmt.Printf("Can't write bytes %s \n", err)
	} else {
		fmt.Printf("Sent %d bytes \n", write)
	}
}
