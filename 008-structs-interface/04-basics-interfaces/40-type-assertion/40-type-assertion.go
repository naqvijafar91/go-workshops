package main

import (
	"fmt"
)

func main() {
	var x interface{} = 7 // x has dynamic type int and value 7
	fmt.Printf("%T\n", x)
	i := x.(int) // i has type int and value 7
	fmt.Printf("%T\n", i)

}
