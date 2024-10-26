package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func createTcpConn(clientId int, wg *sync.WaitGroup) {
	defer wg.Done()
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("\nConnected to server from client:%d", clientId)
	_, err = conn.Write([]byte("Hello"))
	if err != nil {
		fmt.Printf("\nseeing %s from client:%d ", err, clientId)
		return
	}

	time.Sleep(30 * time.Second)

	conn.Close()
}
