package main

import (
	"fmt"
)

func main() {
	myChannel := make(chan string)

	go func() {
		myChannel <- "data"
	}()
	
	// blocking
	msg := <- myChannel
	fmt.Println(msg)
}
