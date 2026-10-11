package main

import "fmt"

func main() {
	initSeq := func(init int) func() int {
		i := init
		return func() int {
			i = i + 1
			return i
		}
	}

	inc := initSeq(5)
	fmt.Println("value=", inc())
	fmt.Println("value=", inc())
	fmt.Println("value=", inc())

	next := initSeq(0)
	fmt.Println("newValue=", next())
}
