package main

import (
	"fmt"
	"log"
	"net"
)

type TCPClient struct {
	laddr string
}

func main() {
	tcpClient := NewTCPClient("0.0.0.0:3333")

	if err := tcpClient.Open(); err != nil {
		log.Fatal("could not open tcpClient", err)
	}
}

func NewTCPClient(laddr string) *TCPClient {
	return &TCPClient{laddr}
}

func (s *TCPClient) Open() error {

	tcpAddr, err := net.ResolveTCPAddr("tcp", s.laddr)
	if err != nil {
		log.Fatal(err)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		log.Fatal(err)
	}

	// Connect to server
	_, err2 := conn.Write([]byte(`Connected`))
	if err != nil {
		log.Fatal(err2)
	}

	var buf [512]byte

	// Read input values
	n, err := conn.Read(buf[0:])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(buf[0:n]))

	return nil
}