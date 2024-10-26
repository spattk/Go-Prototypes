package main

import (
	"fmt"
	"sync"
	"time"
)

type BlockingQueueChannel struct {
	channel chan interface{}
}

func NewBlockingQueueChannel(capacity int) *BlockingQueueChannel {
	q := new(BlockingQueueChannel)
	q.channel = make(chan interface{}, capacity)
	return q
}

func (bq *BlockingQueueChannel) Put(item interface{}) {
	bq.channel <- item
}

func (bq *BlockingQueueChannel) Take() interface{} {
	return <-bq.channel
}
 
func TestConnectionPoolingBufferedChannels(bq *BlockingQueueChannel) {
	var wg sync.WaitGroup
	fmt.Println("Executing DB requests by acquiring DB Connection")
	for i :=1; i <= 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			
			//Grab an available connection to execute the request
			item := bq.Take()
			fmt.Printf("Executing %v DB request using : %v\n", i, item)
			time.Sleep(2000 * time.Millisecond)
			
			//Release the connection after the work is done
			bq.Put(item)
		}()
	}
	wg.Wait()
}