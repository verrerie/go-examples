package main

import "fmt"

func main() {
	cleanup := func() {
		if r := recover(); r != nil {
			fmt.Println("recovered:", r)
		}
	}

	myPanic := func() {
		panic("pb!")
	}

	defer cleanup()

	myPanic()

	fmt.Println("after panic: if you see this message, there is a mistake :D")
}
