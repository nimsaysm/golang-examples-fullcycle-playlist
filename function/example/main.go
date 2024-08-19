package main

import "fmt"

func main() {
	result1 := sum(1, 1)
	result2 := sumAllItems(1, 1, 1, 1, 2)

	fmt.Println(result1)
	fmt.Println(result2)
}

func sum(a, b int) (result int) {
	result = a + b
	return //will take the closest variable
}

// handle x as a array
func sumAllItems(x ...int) int {
	result := 0

	//_ to substitute a key
	for _, value := range x {
		result += value
	}
	return result
}
