package main

import (
	"fmt"
	"hello/math"
)

func main() {
	fmt.Println("Hello, world!")

	resultSum := math.Sum(1, 1)
	// %v -> value of result
	fmt.Printf("%v", resultSum)
}
