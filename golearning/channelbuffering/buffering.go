package main

import (
	"fmt"
	"reflect"
)

func main() {
	s := []int{}

	s = append(s, 1)
	fmt.Println(s)

	q := make([]int, 10, 15)

	q = append(q, 1)
	fmt.Println(q)

	messages := make(chan string, 2)

	fmt.Println(reflect.TypeOf(messages))

	messages <- "buffered"
	messages <- "channel"

	// for i := 0; i < 1; i++ {
	// 	messages <- fmt.Sprintf("%d-message", i)
	// }
	fmt.Println(<-messages)
	fmt.Println(<-messages)

	// Why does commented code gives next mistake?
	//		fatal error: all goroutines are asleep - deadlock!

	// for i := 0; i < 1; i++ {
	// 	<-messages
	// }
}
