package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:7000")
	if err != nil {
		// handle error
	}
	fmt.Println("Client connected to server.")
}