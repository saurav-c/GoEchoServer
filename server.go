package main

import (
	"fmt"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", "localhost:7000")
	if err != nil {
	// handle error
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
		}
		fmt.Println("Added a client...")
	}
}