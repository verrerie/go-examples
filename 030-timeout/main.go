package main

import (
	"fmt"
	"time"
)

func main() {
	work := func(c chan<- string, msg string) {
		time.Sleep(1 * time.Second)
		c <- msg
	}

	receiveWithin := func(c <-chan string, timeout time.Duration, name string) {
		select {
		case res := <-c:
			fmt.Println(name, ":", res)
		case <-time.After(timeout):
			fmt.Println("timeout:", name)
		}
	}

	ch := make(chan string)
	go work(ch, "dev")

	receiveWithin(ch, time.Second/2, "work1")

	ch2 := make(chan string)
	go work(ch2, "plan")

	receiveWithin(ch2, time.Second*2, "work2")

	go work(ch, "sleep")
	receiveWithin(ch, time.Second*2, "work1.2")
}
