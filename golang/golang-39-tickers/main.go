package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(200 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				fmt.Println("work done")
				return
			case <-ticker.C:
				fmt.Println("ticker at", time.Now())
			}
		}
	}()

	time.Sleep(1 * time.Second)
	ticker.Stop()
	done <- true
	fmt.Println("completed")
}
