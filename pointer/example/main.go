package main

import "fmt"

func main() {
	a := 10

	//will handle the memory address of a
	var pointer *int = &a
	//will print the value contained in the address of a (a value)
	fmt.Println(*pointer)

	//changes value via pointer
	*pointer = 50
	fmt.Println(*pointer)
	fmt.Println(pointer)
}
