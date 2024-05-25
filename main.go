package main

import (
	"fmt"
	"net"
	"github.com/ssenthilnathan3/resp"
)

func main() {
	fmt.Println("Listening on port :6379")

	l, err := net.Listen("tcp", ":6379")
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}

	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting: ", err)
		return
	}

	for {
		resp := NewResp(conn)
		value, err := resp.Read()
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(value)

		// ignore request and send back a PONG
		conn.Write([]byte("+OK\r\n"))

	}

}
