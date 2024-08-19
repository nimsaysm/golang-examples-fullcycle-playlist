package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	//if the response takes less than 5 seconds, the request will be canceled
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

	defer cancel()

	request, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080", nil)
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Println(err)
	}

	defer response.Body.Close()

	//prints the request result on the screen
	io.Copy(os.Stdout, response.Body)
}
