package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

func main() {
	//example using http get
	response, error := http.Get("https://google.com.br")

	if error != nil {
		log.Fatal(error.Error())
	}

	fmt.Println(response.Header)

	//example using local func error
	res, err := sum(10, 10)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(res)

	//do not handle the error
	res2, _ := sum(11, 11)

	fmt.Println(res2)
}

func sum(x, y int) (int, error) {
	result := x + y

	if result > 10 {
		return 0, errors.New("the sum is bigger than 10")
	}

	return result, nil
}
