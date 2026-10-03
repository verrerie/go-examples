package main

import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(time.Second / 2)

	<-timer1.C
	fmt.Println("timer1 fired")

	timer2 := time.NewTimer(time.Second / 4)

	go func() {
		<-timer2.C
		fmt.Println("timer2 fired")
	}()

	// repalce / 2 by / 8 would make the .Stop() successful
	time.Sleep(time.Second / 2)
	stopped := timer2.Stop()
	if stopped {
		fmt.Println("timer2 stopped")
	} else {
		fmt.Println("timer2 already stopped")
	}
}
