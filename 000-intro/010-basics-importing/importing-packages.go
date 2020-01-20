package main

import (
	"fmt"

	// if this code will be pushed to github everyone
	// can import in that way:
	"github.com/naqvijafar91/go-workshops/000-intro/010-basics-importing/sub"
	// "./sub" will work too but it's not idiomatic in GO
)

// now we can use some data from our `sub` package
func main() {
	fmt.Println("sub.FirstConstant:\t", sub.FirstConstant)
	fmt.Println("sub.Hot() func call:\t", sub.Hot())
	fmt.Println("sub.SecondConstant:\t", sub.SecondConstant)
}
