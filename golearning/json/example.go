package main

import (
	"encoding/json"
	"fmt"
)

type responseFirst struct {
	Page   int
	Fruits []string
}

type responseSecond struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruites"`
}

func main() {
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))
	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))
	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	resp := responseFirst{
		Page:   12,
		Fruits: []string{"one", "two"},
	}

	res, _ := json.Marshal(resp)
	fmt.Println(string(res))

	respSec := responseSecond{
		Page:   45,
		Fruits: []string{"ttt", "qqq"},
	}
	resSec, _ := json.Marshal(respSec)
	fmt.Println(string(resSec))

}
