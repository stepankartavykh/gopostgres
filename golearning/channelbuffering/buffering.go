package main

import (
	"fmt"
)

func main() {
	messages := make(chan string, 3)

	// fmt.Println(reflect.TypeOf(messages))

	messages <- "message 1"
	messages <- "message 2"
	messages <- "message 3"

	// for i := range 1 {
	// 	messages <- fmt.Sprintf("%d-message", i)
	// }
	fmt.Println(<-messages)
	fmt.Println(<-messages)
	fmt.Println(<-messages)

	// Why does commented code gives next mistake?
	//		fatal error: all goroutines are asleep - deadlock!

	// for i := 0; i < 1; i++ {
	// 	<-messages
	// }
}
