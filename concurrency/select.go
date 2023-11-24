package main

import (
	"fmt"
)

func main() {
	chan1 := make(chan string)
	chan2 := make(chan string)

	go func() {
		chan1 <- "learning GO"
	} ()

	go func() {
		chan2 <- "learning concurrency in GO"
	} ()

	// block until it receives message from a channel
	// if multiple at sametime then selects at random
	select {
		case msgFromChan1 := <- chan1:
			fmt.Println(msgFromChan1)
		case msgFromChan2 := <- chan2:
			fmt.Println(msgFromChan2)
	}
}
