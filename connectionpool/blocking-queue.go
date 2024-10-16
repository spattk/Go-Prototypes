package main

import (
	"fmt"
	"sync"
	"time"
	"strconv"
)

var dbConnCount = 0

type BlockingQueue struct {
	m        sync.Mutex
	c        sync.Cond
	data     []interface{}
	capacity int
}

func NewBlockingQueue(capacity int) *BlockingQueue {
	q := new(BlockingQueue)
	q.c = sync.Cond{L: &q.m}
	q.capacity = capacity
	return q
}

func (q *BlockingQueue) Put(item interface{}) {
	q.c.L.Lock()
	defer q.c.L.Unlock()

	for q.isFull() {
		fmt.Println("Full, waiting...")
		q.c.Wait()
	}

	q.data = append(q.data, item)
	q.c.Signal()
}

func (q *BlockingQueue) isFull() bool {
	return len(q.data) == q.capacity
}

func (q *BlockingQueue) isEmpty() bool {
	return len(q.data) == 0
}

func (q *BlockingQueue) Take() interface{} {
	q.c.L.Lock()
	defer q.c.L.Unlock()

	for q.isEmpty() {
		fmt.Println("No connections available, waiting for a connection ...")
		q.c.Wait()
	}

	result := q.data[0]
	q.data = q.data[1:]
	q.c.Signal()
	return result
}

func TestConnectionPoolingBlockingQueue(bq *BlockingQueue) {
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

func getDBConnection() string {
	dbConnCount++;
	return "db-conn-"+strconv.Itoa(dbConnCount);
}


