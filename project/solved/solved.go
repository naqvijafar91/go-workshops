package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func slowRead() string {
	time.Sleep(1 * time.Second)
	return "Response from Sleep Read"
}

func slowWrite(value string) {
	time.Sleep(1 * time.Second)
	fmt.Println(value)
}

func main() {
	start := time.Now()
	// Modify the code below this line

	times := 500
	wg := &sync.WaitGroup{}
	wg.Add(times)
	readCh := make(chan string)
	for i := 0; i < times; i++ {
		go func() {
			readCh <- slowRead()
		}()
	}
	for i := 0; i < times; i++ {
		go func() {
			slowWrite(<-readCh)
			wg.Done()
		}()

	}

	wg.Wait()

	// Modify the code above this line
	elapsed := time.Since(start)
	log.Printf("%d Tasks took %s", times,elapsed)
}
