package main

import (
	"fmt"

	"github.com/naqvijafar91/go-workshops/008-structs-interface/00-basics-structs-defining/image"
)

func main() {
	img := image.NewImage(1888)
	// The below statement will fail as size is a private field of the struct
	// fmt.Printf("%+v", img.size)
	fmt.Printf("%+v", img.GetSize())
}
