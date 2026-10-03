package main

import "fmt"

func main() {
	ch := make(chan string, 2)

	send := func(msg string) {
		fmt.Println("sending message:", msg)
		ch <- msg
	}
	go send("hello")
	go send("hi")

	fmt.Println("waiting for message")
	n := 0
	for msg := range ch {
		fmt.Println("received:", msg)
		n = n + 1
		if n >= 2 {
			fmt.Println("all messages received")
			return
		}

	}
}
