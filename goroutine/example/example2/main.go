package main

import (
	"fmt"
)

// main = thread 1
func main() {
	//channel (chan) between thread 1 and thread 2
	hello := make(chan string)

	// thread 2:
	go func() {
		hello <- "Hello world"
	}()

	result := <-hello
	fmt.Println(result)
}
