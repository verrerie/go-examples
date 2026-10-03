package main

import (
	"fmt"
	"time"
)

func main() {
	basicLimiter(5)
	burstLimiter(4)
}

func basicLimiter(num int) {
	requests := make(chan int, num)
	for i := range num {
		requests <- i
	}
	close(requests)

	limiter := time.Tick(100 * time.Millisecond)
	for t := range requests {
		<-limiter
		fmt.Println("sending:", t, time.Now())
	}
}

func burstLimiter(num int) {
	burstLimit := make(chan time.Time, num)

	for range num {
		burstLimit <- time.Now()
	}

	go func() {
		for t := range time.Tick(100 * time.Millisecond) {
			burstLimit <- t
		}
	}()

	burstRequests := make(chan int, num+4)
	for i := range num + 4 {
		burstRequests <- i
	}
	close(burstRequests)
	for req := range burstRequests {
		<-burstLimit
		fmt.Println("sending", req, time.Now())
	}
}
