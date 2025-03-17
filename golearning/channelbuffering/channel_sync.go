package main

import (
	"fmt"
	"os"
	"time"
)

func decorator(function func()) {
	function()
}

func worker(done chan bool) {
	fmt.Println(os.Getpid())
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func example() {
	fmt.Println("example function")
}

func main() {
	fmt.Println(os.Getpid())
	done := make(chan bool, 1)
	go worker(done)

	<-done
	decorator(example)
}
