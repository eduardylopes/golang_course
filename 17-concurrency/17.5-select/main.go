package main

import (
	"fmt"
	"time"
)

func main() {
	channel1, channel2 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			channel1 <- "Channel #1"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			channel2 <- "Channel #2"
		}
	}()

	for {
		// Used only for concurrency, if message exists print the message, in this case the "chan" not block the scope.
		select {
		case channel1Message := <-channel1:
			fmt.Println(channel1Message)
		case channel2Message := <-channel2:
			fmt.Println(channel2Message)
		}
	}
}
