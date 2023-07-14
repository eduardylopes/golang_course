package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan string)
	go write("Goroutine #1", channel)

	fmt.Println("This will be executed normally")

	for mensagem := range channel {
		fmt.Println(mensagem)
	}

	fmt.Println("Fim do programa!")
}

func write(text string, channel chan string) {
	for i := 0; i < 5; i++ {
		channel <- text
		time.Sleep(time.Second)
	}

	close(channel)
}
