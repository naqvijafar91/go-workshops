package main

import (
	"fmt"
)

func sequencer() func() int {
	i := 0

	return func() int {
		i++
		return i
	}
}
func main() {
	// nextInt has its own data
	nextInt := sequencer()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	//newInt does not have nextInt's data
	newInts := sequencer()
	fmt.Println(newInts())
	fmt.Println(newInts())

}
