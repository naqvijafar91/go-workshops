package main

import (  
    "fmt"
    "time"
)

func main() {  
    data := []string{"one","two","three"}
	// What will be the output?
    for _,v := range data {
        go func() {
            fmt.Println(v)
        }()
    }

    time.Sleep(3 * time.Second)
}