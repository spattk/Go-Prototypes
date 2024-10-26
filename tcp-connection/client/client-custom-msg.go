package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func createTcpConnCustomMsg() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Connected to server")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a message to send to server")
	str, err := reader.ReadString('\n')
	_, err = conn.Write([]byte(str))
	if err != nil {
		fmt.Println(err)
		return
	}

	conn.Close()
}
