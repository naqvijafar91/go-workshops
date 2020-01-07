// to jest dziwne :/
package main

import "fmt"

type data struct {
	Name string
}

func main() {
	n := make(map[string]data)
	n["x"] = data{"one"}
	fmt.Println(n["x"])
	n["x"].Name = "three"
	fmt.Println(n["x"])
}
