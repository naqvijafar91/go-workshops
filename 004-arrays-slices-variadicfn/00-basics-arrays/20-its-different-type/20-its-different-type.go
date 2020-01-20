package main

import "fmt"

func main() {
	var five [5]int
	four := [4]int{10, 20, 30, 40}
	// The following statement will not work as five is an array of 5 elements and four is for 4 elements,
	// so they are a different type
	five = four

	fmt.Println(four)
	fmt.Println(five)
}
