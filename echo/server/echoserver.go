package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net"
)

//Sample echo server
//Respond to the caller with the same buffer
func main() {

	var network string
	var addr string
	flag.StringVar(&network, "n", "tcp", "network type [tcp/unix]")
	flag.StringVar(&addr, "e", ":4040", "service endpoint [addr or path]")

	listen, err := net.Listen(network, addr)
	if err != nil {
		panic(err)
	}

	defer listen.Close()

	fmt.Printf("Listening on network %s and address %s \n", listen.Addr().Network(), listen.Addr())

	for {
		accept, err := listen.Accept()
		if err != nil {
			panic(err)
		}
		fmt.Printf("Accepted connection from %s \n", accept.RemoteAddr())

		go handleConnection(accept)
	}
}

func handleConnection(conn net.Conn) {

	tmp, err := ioutil.ReadAll(conn)
	if err != nil {
		fmt.Println("Error reading data")
	} else {
		fmt.Printf("Read %d bytes -> %s \n", len(tmp), tmp)
	}

	fmt.Println("Responding ... ")
	write, err := conn.Write(tmp)
	if err != nil {
		fmt.Printf("Error writing response %s \n", err)
	} else {
		fmt.Printf("Response written %d \n", write)
	}
}
