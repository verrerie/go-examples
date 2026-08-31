package main

import "fmt"

func main() {
	ping := func(p chan<- string, msg string) {
		fmt.Println("sending:", msg)
		p <- msg
	}

	pong := func(p <-chan string, o chan<- string) {
		msg := <-p
		fmt.Println("received:", msg)
		o <- msg
		fmt.Println("redirected:", msg)
		fmt.Println()
	}

	in := make(chan string, 1)
	out := make(chan string, 1)

	ping(in, "hello world")
	pong(in, out)
	pong(out, in)
	pong(in, out)

	last := out

	fmt.Println("last message:", <-last)
}
