
package main

import (
	"fmt"
)

func main() {

	warehouse := map[string]int{
		"lizak":   100,
		"guma":    1000,
		"arbuz":   12,
		"cytryna": 22,
	}

	for i := 0; i < 5; i++ {
		fmt.Println("Iteration", i)
		fmt.Println("---------------")
		for k, v := range warehouse {
			fmt.Println(k, v)
		}
		fmt.Println("---------------")
	}
}

// Takeaway: Never depend on the order when iterating over a map, if you want to iterate in order, then sort keys and use them