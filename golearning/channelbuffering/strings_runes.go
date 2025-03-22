package main

import (
	"fmt"
	"unicode/utf8"
)

func examineRune(r rune) {
	if r == 't' {
		fmt.Println("found teee")
	} else if r == 'ส' {
		fmt.Println("found suaa")
	}
}

func main() {
	const sOne = "สวัสดี and some english letterss"

	fmt.Println(len(sOne))

	const sSecond = "สวัสดี"

	fmt.Println(len(sSecond))

	for i := 0; i < len(sSecond); i++ {
		fmt.Printf("%x ", sSecond[i])
	}
	fmt.Println("Rune count:", utf8.RuneCountInString(sSecond))

	for idx, runeValue := range sSecond {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

	fmt.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(sSecond); i += w {
		runeValue, width := utf8.DecodeRuneInString(sSecond[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width
		examineRune(runeValue)
	}

	const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"
	fmt.Println(sample)
	for i := 0; i < len(sample); i++ {
		fmt.Printf("%x ", sample[i])
	}
	fmt.Printf("%x\n", sample)
	fmt.Printf("% x\n", sample)

	fmt.Printf("%q\n", sample)
	fmt.Printf("%+q\n", sample)

	const rawString = `⌘`
	fmt.Printf("plain string: ")
	fmt.Printf("%s", rawString)
	fmt.Printf("\n")

	fmt.Printf("quoted string: ")
	fmt.Printf("%+q", rawString)
	fmt.Printf("\n")

	fmt.Printf("hex bytes: ")
	for i := 0; i < len(rawString); i++ {
		fmt.Printf("%x ", rawString[i])
	}
	fmt.Printf("\n")
}
