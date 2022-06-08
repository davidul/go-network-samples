package main

import (
	"fmt"
	"net"
)

func main() {
	listen, err := net.Listen("tcp4", "0.0.0.0:9000")
	if err != nil {
		panic(err)
	}

	defer listen.Close()

	for {
		accept, err := listen.Accept()
		if err != nil {
			panic(err)
		}

		fmt.Println("Connected to ", accept.RemoteAddr())

		go handleConnection(accept)

	}

	listen.Close()
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	bytes := make([]byte, 1024)
	readBytes, err := conn.Read(bytes)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Read %v bytes", readBytes)
	fmt.Println(bytes[0:readBytes])

}
