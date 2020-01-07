package main

import (
	"flag"
	"fmt"
)

// single flag of string type
var logLevel = flag.String("log_level", "warning", "Sets logging level")

func main() {
	// we need to call Parse(), after that flags will be parsed
	flag.Parse()
	fmt.Println("Current log level:", *logLevel)
}
