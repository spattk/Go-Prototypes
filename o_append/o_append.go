package main

import (
	"fmt"
	"io"
	"os"
	"sync"
)

func getMsgFromIdx(idx int) string {

	newIdx := idx % 5
	messages := map[int]string{
		0: "OOO\n",
		1: "AAA\n",
		2: "BBB\n",
		3: "CCC\n",
		4: "DDD\n",
	}

	return messages[newIdx]
}

func o_append(idx int, wg *sync.WaitGroup) {
	defer wg.Done()
	msg := getMsgFromIdx(idx)
	fd, err := os.OpenFile("log_oappend.txt", os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer fd.Close()

	fd.WriteString(msg)
}

func no_o_append(idx int, wg *sync.WaitGroup) {
	defer wg.Done()
	msg := getMsgFromIdx(idx)
	fd, err := os.OpenFile("log_no_oappend.txt", os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer fd.Close()

	fd.Seek(0, io.SeekEnd)
	fd.WriteString(msg)
}

func main() {
	println("starting...")
	var wg1 sync.WaitGroup
	var wg2 sync.WaitGroup
	
	for i := 0; i < 10; i++ {
		wg1.Add(1)
		wg2.Add(1)
		go o_append(i, &wg1);
		go no_o_append(i, &wg2);
	}
	
	wg1.Wait()
	wg2.Wait()
	println("done...")
}