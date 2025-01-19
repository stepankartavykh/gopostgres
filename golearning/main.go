package main

import "fmt"

func firstWorker() {
	fmt.Println("Test channels to send messages in two goroutines.")
}

func secondWorker() {
	fmt.Println("Test channels to send messages in two goroutines.")
}

func main() {
	firstWorker()
	secondWorker()
	messages := make(chan string)

	go func() { messages <- "ping" }()

	msg := <-messages
	fmt.Println(msg)
}
