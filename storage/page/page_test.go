package page

import (
	"fmt"
	"testing"
)

func TestPageCreation(t *testing.T) {
	fmt.Println("page is created...")
}

func TestTypeDefineFunction(t *testing.T) {
	testStruct := pageLayout{
		Header:        nil,
		Items:         nil,
		PageTotalSize: 0,
	}
	typ := DefineType(testStruct)
	fmt.Println(typ)
}

func TestProcessIdentities(t *testing.T) {
	iter := func() string {
		fmt.Println(t.Deadline())
		return "test"
	}
	iter()
}
