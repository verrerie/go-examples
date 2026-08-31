package main

import (
	"fmt"
	"time"
)

func work(job <-chan string, result chan<- string, name string) {
	fmt.Println("calling", name)
	for j := range job {
		fmt.Println(name, "is fighting against", j)
		time.Sleep(100 * time.Millisecond)
		fmt.Println(name, "won,", j, "defeated")
		result <- name + " beated " + j
	}
}

func main() {
	const numJobs = 5
	job := make(chan string)
	result := make(chan string)

	for w := range numJobs {
		go work(job, result, fmt.Sprintf("superman%v", w))
	}

	for j := range numJobs {
		job <- fmt.Sprintf("bad-guy%v", j)
	}
	close(job)

	for range numJobs {
		fmt.Println("Justice:", <-result)
	}
}
