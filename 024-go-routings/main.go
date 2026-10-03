package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := range 5 {
		fmt.Println(from, i)
	}
}

func main() {
	f("direct")

	go f("go routine")
	go f("go routine2")

	go func(msg string) {
		fmt.Println(msg)
	}("single message here!")

	fmt.Println("start sleep")
	time.Sleep(time.Second)
	fmt.Println("done")
}
