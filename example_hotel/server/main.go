package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":8080", nil)
}

func home(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	log.Println("Starts request")
	defer log.Println("Ends request")

	select {
	case <-time.After(time.Second * 5):
		log.Println(writer, "Successful request")
		writer.Write([]byte("Successful request"))
	case <-ctx.Done():
		log.Println("Canceled request")
		http.Error(writer, "Canceled request", http.StatusRequestTimeout)
	}
}
