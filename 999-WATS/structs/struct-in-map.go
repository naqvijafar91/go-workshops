package main

import "fmt"

type data struct {
	Name string
}

func main() {
	n := make(map[string]data)
	n["x"] = data{"one"}
	fmt.Println(n["x"])
	n["x"].Name = "three" // that's weird : shorturl.at/aorz5
	fmt.Println(n["x"])
}
