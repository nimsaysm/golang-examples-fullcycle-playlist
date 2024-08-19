package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	//Background -> main context
	ctx := context.Background()
	//cancel -> returns a cancel function
	ctx, cancel := context.WithCancel(ctx)

	//will wait all executions for execute cancel
	defer cancel()

	go func() {
		time.Sleep(time.Second * 4) //action that cancels the context
		cancel()
	}()

	bookHotel(ctx)
}

func bookHotel(ctx context.Context) {
	//select -> executes when something (channel) worked
	select {
	case <-ctx.Done():
		fmt.Println("Exceeded booking time.")
	case <-time.After(time.Second * 5):
		fmt.Println("Successful booking!")
	}
}
