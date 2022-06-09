package main

import (
	"bytes"
	"encoding/gob"
	"flag"
	"fmt"
	"net"
	"os"
)

type Message struct {
	Cmd     string
	Payload string
}

func main() {

	var addr string
	var network string
	flag.StringVar(&network, "n", "tcp", "network protocol [tcp/unix]")
	flag.StringVar(&addr, "e", ":4040", "service endpoint [ip address or socket path]")

	switch network {
	case "tcp", "tcp4", "tcp6", "unix":
	default:
		fmt.Println("Unsupported protocol")
		os.Exit(1)
	}

	listen, err := net.Listen(network, addr)
	if err != nil {
		panic(err)
	}

	defer listen.Close()

	fmt.Printf("Listenting at %s %s \n", network, addr)

	for {
		accept, err := listen.Accept()
		if err != nil {
			panic(err)
		}

		fmt.Println("Connected to ", accept.RemoteAddr())

		go handleConnection(accept)

	}
}

func handleConnection(conn net.Conn) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			panic(err)
		} else {
			fmt.Println("Disconnected from:", conn.RemoteAddr())
		}
	}(conn)

	tmpBuf := make([]byte, 1024)
	readBytes, err := conn.Read(tmpBuf)
	if err != nil {
		fmt.Println("Error reading: ", err)
	} else {
		fmt.Printf("Read %v tmpBuf \n", readBytes)
		//fmt.Println(string(tmpBuf[0:readBytes]))
	}

	buffer := bytes.NewBuffer(tmpBuf)

	tmpMessage := new(Message)
	decoder := gob.NewDecoder(buffer)
	err = decoder.Decode(tmpMessage)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Message: %v %v \n", tmpMessage.Cmd, tmpMessage.Payload)
}
