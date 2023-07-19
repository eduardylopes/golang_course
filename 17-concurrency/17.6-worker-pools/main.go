package main

import (
	"fmt"
	"time"
)

var fibTerm = 45
var fibCache = make(map[int]int, fibTerm)

func main() {
	tasks := make(chan int, fibTerm)
	results := make(chan int, fibTerm)

	go worker(tasks, results)
	go worker(tasks, results)
	go worker(tasks, results)

	for i := 0; i < fibTerm; i++ {
		tasks <- i
	}
	close(tasks)

	start := time.Now()
	for i := 0; i < fibTerm; i++ {
		result := <-results
		fmt.Println(result)
	}

	fmt.Printf("Execution time: %v\n", time.Since(start))
}

func worker(tasks <-chan int, results chan<- int) {
	for number := range tasks {
		result := fibonacci(number)
		fibCache[number] = result
		results <- result
	}
}

func fibonacci(number int) int {
	if value, ok := fibCache[number]; ok {
		return value
	}

	if number < 2 {
		return number
	}

	return fibonacci(number-2) + fibonacci(number-1)
}
