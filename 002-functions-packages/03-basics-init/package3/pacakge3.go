package package3

import (
	"fmt"
	"io"

	"go-workshops/002-functions-packages/03-basics-init/package1"
	"go-workshops/002-functions-packages/03-basics-init/package2"
	"go-workshops/002-functions-packages/03-basics-init/package4"
)

func init() {
	fmt.Println("Initialize package3!!!")
}

func Write(writer io.Writer) {
	package1.Write(writer)
	package2.Write(writer)
	package4.Write(writer)
}
