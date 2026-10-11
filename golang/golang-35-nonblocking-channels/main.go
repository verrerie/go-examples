package main

import "fmt"

func main() {
	// with buffer size = 1, we can send and receive a message synchronously
	// msg := make(chan string, 1)
	msg := make(chan string)
	sig := make(chan bool)

	select {
	case m := <-msg:
		fmt.Println("message received:", m)
	default:
		fmt.Println("no message received")

	}

	select {
	case msg <- "hi":
		fmt.Println("message sent")
	default:
		fmt.Println("no message sent")
	}

	select {
	case m := <-msg:
		fmt.Println("message received:", m)
	case s := <-sig:
		fmt.Println("signal received:", s)
	default:
		fmt.Println("nothing happened")
	}
}
