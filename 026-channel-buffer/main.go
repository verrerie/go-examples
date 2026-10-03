package main

import "fmt"

func main() {
	ch := make(chan int, 2)

	ch <- 1
	ch <- 2
	// this will cause deadlock, because there is only two slots in the buffer and no one reads them yet
	// ch <- 2

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
