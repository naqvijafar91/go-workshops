package main

import (
	"fmt"
	"os"

	"go-workshops/002-functions-packages/03-basics-init/package1"
	"go-workshops/002-functions-packages/03-basics-init/package2"
	"go-workshops/002-functions-packages/03-basics-init/package3"
)

func main() {
	fmt.Println("Start")
	package1.Write(os.Stdout)
	package2.Write(os.Stdout)

	fmt.Println("Writing all inside package3")
	package3.Write(os.Stdout)
}
