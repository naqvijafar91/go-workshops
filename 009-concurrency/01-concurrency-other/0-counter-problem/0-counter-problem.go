// there is an access problem in "Concurent Programming"
// to variables, what if 2 gorutines hit at the same time
// or save the value? which will be saved.
// A simple example that illustrates the increment of values
package main

import (
	"fmt"
	// "runtime"
	"time"
)

type intCounter int64

func (c *intCounter) Add(x int64) {
	*c += intCounter(x)
}

func (c *intCounter) Value() (x int64) {
	return int64(*c)
}

func main() {
	counter := intCounter(0)

	for i := 0; i < 100; i++ {
		go func(no int) {
			for i := 0; i < 10000; i++ {
				counter.Add(1)
			}
		}(i)
	}

	time.Sleep(2 * time.Second)
	fmt.Println(counter.Value())

}

// Data races - this is a typical example
// to detect it automatically makes it available
// switch -race

// go run -race 0-counter-problem.go

// https://golang.org/doc/articles/race_detector.html
