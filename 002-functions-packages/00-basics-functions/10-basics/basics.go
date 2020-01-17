package main

import (
	"fmt"
)

//   / func name
//   |      / parameters
//   |      |                  / return type
//   |      |                 |
func getBoo(boosCount int) string {
	return fmt.Sprintf("There are %d boo(s)", boosCount)
}

// There can be only one main() function in package
// Official spec:The main package must have package name main and
// declare a function main that takes no arguments and returns no value.
// Program execution begins by initializing the main package and then invoking the function main
// When that function invocation returns, the program exits.
// It does not wait for other (non-main) goroutines to complete.
func main() {
	println(getBoo(3))
}
