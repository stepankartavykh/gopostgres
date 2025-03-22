package main

import "fmt"

func main() {
	const sample = "sample"

	for _, symbol := range sample {
		fmt.Printf("%v %d %x\n", symbol, symbol, symbol)
	}
}
