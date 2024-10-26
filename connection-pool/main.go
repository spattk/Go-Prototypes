package main

import (
	"fmt"
)

func main() {
	blockingQ := false
	poolSize := 10
	if blockingQ {
		fmt.Println("USING MUTEX TO IMPLEMENT BLOCKING QUEUE")
		bq := NewBlockingQueue(poolSize)
		fmt.Println("Initiating Connection pooling...")
		for i := 1; i <= poolSize; i++ {
			bq.Put(getDBConnection())
		}
		fmt.Println("10 initial connections established.")
		TestConnectionPoolingBlockingQueue(bq)
	} else {
		fmt.Println("USING BUFFERED CHANNEL TO IMPLEMENT BLOCKING QUEUE")
		bq := NewBlockingQueueChannel(poolSize)
		fmt.Println("Initiating Connection pooling...")
		for i := 1; i <= poolSize; i++ {
			bq.Put(getDBConnection())
		}
		fmt.Println("10 initial connections established.")
		TestConnectionPoolingBufferedChannels(bq)
	}
	
}