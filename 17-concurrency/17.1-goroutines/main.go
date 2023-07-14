package main

import (
	"fmt"
	"time"
)

func main() {
	go write("Goroutine #1") // goroutine
	write("Hello world!")
}

func write(text string) {
	for {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
