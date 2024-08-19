package main

import (
	"fmt"
	"time"
)

func counter(counterType string) {
	for i := 0; i < 5; i++ {
		fmt.Println(counterType, i)
		time.Sleep(time.Second)
	}
}

func main() {
	// a and b will be executed a the same time
	go counter("a")
	go counter("b")
	time.Sleep(time.Second * 10)
}
