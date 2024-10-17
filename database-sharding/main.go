package main

import (
	"fmt"
	"sync"
	"math/rand"
    "time"
)


func main() {

	userIdMap := make(map[int][]int)
	fmt.Println("Adding Users to DB...")
	for i :=1; i <= 100; i++ {
		addUserDetails(i, userIdMap)
	}
	fmt.Println("Adding Users completed")

	reqSize := 20
	var wg sync.WaitGroup
	rand.New(rand.NewSource(time.Now().UnixNano()))


	fmt.Println("Querying users from the DB...")
	for i := 1; i <= reqSize; i++ {
		userId := rand.Intn(100)

		wg.Add(1)
		go getUserDetails(&wg, userId, userIdMap)
	}

	// wait for all the go-routines to complete
	wg.Wait()
}
