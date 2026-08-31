package main

import (
	"fmt"
	"time"
)

func main() {
	worker := func(name string, complete chan bool) {
		fmt.Println("workder", name, "is working")
		time.Sleep(time.Second / 2)
		fmt.Println(name, "done")
		complete <- true
	}

	done := make(chan bool)

	go worker("bot1", done)
	<-done
	// without the buffer size 1, worker() will block infinitely
	ok := make(chan bool, 1)
	worker("bot2", ok)
	<-ok
}
