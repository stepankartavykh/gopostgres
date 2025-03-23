package main

import "fmt"

type SomeCode int

const (
	first  SomeCode = 0
	second SomeCode = 1
	third  SomeCode = 2
)

func returnErrorCode() SomeCode {
	return 5
}

func main() {
	b := byte('\n')

	fmt.Println(b)

	fmt.Println(b + 10)

	fmt.Println(first)

	fmt.Println(returnErrorCode())

}
