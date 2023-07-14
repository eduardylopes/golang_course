package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	channel1 := write("Hello world!")
	channel2 := write("Programming in Go!")

	multiplexedChannel := multiplex(channel1, channel2)

	for i := 0; i < 10; i++ {
		fmt.Println(<-multiplexedChannel)
	}
}

func multiplex(inputChannel1, inputChannel2 <-chan string) <-chan string {
	outputChannel1 := make(chan string)

	go func() {
		for {
			select {
			case message := <-inputChannel1:
				outputChannel1 <- message
			case message := <-inputChannel2:
				outputChannel1 <- message
			}
		}
	}()

	return outputChannel1
}

func write(text string) <-chan string {
	channel := make(chan string)

	go func() {
		for {
			channel <- fmt.Sprintf("Received value: %s", text)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		}
	}()

	return channel
}
