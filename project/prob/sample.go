package main

import (
	"fmt"
	"log"
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
	
	times := 5
	for i := 0; i < times; i++ {
		go slowWrite(slowRead())
	}

	// Modify the code above this line
	elapsed := time.Since(start)
	log.Printf("Task took %s", elapsed)
}
