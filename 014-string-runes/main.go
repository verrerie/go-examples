package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	const s = "สวัสดี"
	fmt.Println("len=", len(s))

	for r := range s {
		fmt.Printf("%x ", r)
	}
	fmt.Println()
	fmt.Println("runes count:", utf8.RuneCountInString(s))

	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}
}
