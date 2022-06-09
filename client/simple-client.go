package main

import (
	"bytes"
	"encoding/gob"
	"flag"
	"net"
)

type Message struct {
	Cmd     string
	Payload string
}

func main() {

	msgBuffer := new(bytes.Buffer)
	encoder := gob.NewEncoder(msgBuffer)

	var addr string
	var network string
	flag.StringVar(&addr, "e", "localhost:4040", "service address endpoint")
	flag.StringVar(&network, "n", "tcp", "network protocol to use")
	flag.Parse()

	text := flag.Arg(0)

	conn, err := net.Dial("tcp", "localhost:4040")
	if err != nil {
		panic(err)
	}

	defer conn.Close()
	message := Message{Cmd: "HELLO", Payload: text}

	err = encoder.Encode(message)
	if err != nil {
		panic(err)
	}
	_, err = conn.Write(msgBuffer.Bytes())
	//_, err = conn.Write([]byte(text))
	if err != nil {
		panic(err)
	}
}
