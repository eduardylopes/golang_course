package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup

	// Add 4 goroutines
	waitGroup.Add(4)

	go func() {
		write("Goroutine #1")
		waitGroup.Done() // -1
	}()

	go func() {
		write("Goroutine #2")
		waitGroup.Done() // -1
	}()

	go func() {
		write("Goroutine #3")
		waitGroup.Done() // -1
	}()

	go func() {
		write("Goroutine #4")
		waitGroup.Done() // -1
	}()

	// Wait for all goroutines to done
	waitGroup.Wait()
}

func write(text string) {
	for i := 0; i < 5; i++ {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
