package main

import (
	"fmt"

	"github.com/naqvijafar91/go-workshops/060-basics-structs-defining/image"
)

func main() {
	img := image.NewImage(1888)

	fmt.Printf("%+v", img.GetSize())
}
