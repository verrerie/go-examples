package main

import "fmt"

func main() {
	const numberMessages = 5
	queue := make(chan string, numberMessages)

	for i := range numberMessages {
		queue <- fmt.Sprintf("idea %v", i)
	}
	close(queue)

	for m := range queue {
		fmt.Println("got:", m)
	}

	m, more := <-queue
	fmt.Println("after closing, we got:", m, more)
}
