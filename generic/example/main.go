package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type Number interface {
	~int | ~int64 | ~float64
}

type MyNumber int

func GenericSum[T Number](m map[string]T) T {
	var sum T
	for _, v := range m {
		sum += v
	}
	return sum
}

func IsEqual[T comparable](num1, num2 T) bool {
	return num1 == num2
}

func Max[T constraints.Ordered](num1, num2 T) T {
	if num1 > num2 {
		return num1
	}
	return num2
}

func main() {
	var x, y MyNumber
	x = 1
	y = 1

	fmt.Println(GenericSum(map[string]MyNumber{"x": x, "y": y}))

	fmt.Println(GenericSum(map[string]int64{"x": 1, "y": 2, "z": 1}))
	fmt.Println(GenericSum(map[string]float64{"x": 1.5, "y": 2.5, "z": 1.5}))

	fmt.Println(IsEqual(1, 1))
	fmt.Println(Max(7, 6))
}
