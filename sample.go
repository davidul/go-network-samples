package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"strings"
)

type Pxy struct{}

func (p *Pxy) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	fmt.Printf("Received request %s %s %s\n", req.Method, req.Host, req.RemoteAddr)
	transport := http.DefaultTransport
	// step 1
	outReq := new(http.Request)
	*outReq = *req // this only does shallow copies of maps
	if clientIP, _, err := net.SplitHostPort(req.RemoteAddr); err == nil {
		if prior, ok := outReq.Header["X-Forwarded-For"]; ok {
			clientIP = strings.Join(prior, ", ") + ", " + clientIP
		}
		outReq.Header.Set("X-Forwarded-For", clientIP)

	}
	// step 2
	res, err := transport.RoundTrip(outReq)
	if err != nil {
		rw.WriteHeader(http.StatusBadGateway)
		return
	}
	// step 3
	for key, value := range res.Header {
		for _, v := range value {
			rw.Header().Add(key, v)
		}
	}
	rw.WriteHeader(res.StatusCode)
	_, err = io.Copy(rw, res.Body)
	if err != nil {
		log.Println(err)
		return
	}

	err = res.Body.Close()
	if err != nil {
		log.Println(err)
		return
	}
}

func tcp() {
	listen, err := net.Listen("tcp4", "0.0.0.0:9090")
	if err != nil {
		panic(err)
	}
	defer listen.Close()

	accept, err := listen.Accept()
	if err != nil {
		panic(err)
	}

	bytes := make([]byte, 1024)
	read, err := accept.Read(bytes)
	if err != nil {
		panic(err)
	}

	fmt.Println(read)
	accept.Close()
}

func main() {
	fmt.Println("Serve on :8080")
	//http.Handle("/", &Pxy{})
	//http.ListenAndServe("0.0.0.0:8080", nil)
	tcp()
}
