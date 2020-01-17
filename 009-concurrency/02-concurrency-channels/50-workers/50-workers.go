package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	// Each worker will endlessly seek out jobs and execute them as soon as it is done with a single job
	// range loop will end when close is called on the jobs channel
	for j := range jobs {
		time.Sleep(time.Second)
		fmt.Println("worker", id, "processing job", j)
		results <- j * j
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)
	// We create 3 workers to take up as many jobs as required
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}
	// Add jobs
	for j := 1; j <= 9; j++ {
		jobs <- j
	}
	// Close the jobs channel signalling no more jobs will come
	close(jobs) 

	for a := 1; a <= 9; a++ {
		fmt.Println("result: ", a, " = ", <-results)
	}
}
