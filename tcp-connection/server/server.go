package main

import (
    "fmt"
    "net"
)

func main() {
	//Setting up a tcp listener at port 8080
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}

	//Accepting connections
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		// handle connection
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 1024)
	nb, err := conn.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("\nRead %d bytes with text %s", nb, buf)
}