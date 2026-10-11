package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	go func() {
		for i := range 10 {
			c1 <- i
			time.Sleep(time.Second / 10)
		}
	}()

	go func() {
		for i := range 10 {
			c2 <- i
			time.Sleep(time.Second / 5)
		}
	}()

	run(c1, c2)
}

func run(c1 chan int, c2 chan int) {
	var m1, m2 int
	for {
		select {
		case m1 = <-c1:
			fmt.Println(m1, "<-channel-1")
		case m2 = <-c2:
			fmt.Println(m2, "<-channel-2")
		}
		fmt.Println("product:", m1, "x", m2, "=", m1*m2)
		if m1*m2 > 50 {
			fmt.Println("bye")
			return
		}
	}
}
