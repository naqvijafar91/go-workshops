package main

import "fmt"

var (
	illKillYou int
)

func init() {
	// Comment the below line to see the magic
	illKillYou := 1
	illKillYou++
	illKillYou++
}

func main() {
	fmt.Println(illKillYou)

	other := 10

	{
		other := 999
		fmt.Println("inside block", other)
	}

	fmt.Println("outside block", other)

}

// Take away: Shadow variables are bad and produce buggy code
// go tool vet -shadow 10-function-scope.go
