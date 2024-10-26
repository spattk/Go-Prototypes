package main

import (
	"sync"
)

func main() {
	//single stdin client msg connection
	// createTcpConnectionCustom()

	//20k connections
	var wg sync.WaitGroup
	for i := 0; i < 20000; i++ {
		wg.Add(1)
		go createTcpConn(i, &wg)
	}

	wg.Wait()

}
