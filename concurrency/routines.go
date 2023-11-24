package main

import (
	"fmt"
	"time"
)

func randomFunc(x int) {
	fmt.Printf("Hello %v \n", x)
}

func main() {
	go randomFunc(1)
	go randomFunc(2)

	time.Sleep(2 * time.Second)

	fmt.Println("Bye!")
}
