package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan uint)
	go func() {
		time.Sleep(time.Second)
		channel <- 5
	}()

	// wait until there is a value in the channel
	fmt.Println(<-channel)
}
