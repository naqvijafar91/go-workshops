package main

import (
	"fmt"
)

func main() {
	a := struct {
		ID   int
		Name string
	}{1, "Placek"}

	fmt.Println(a)
}
