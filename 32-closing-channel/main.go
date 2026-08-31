package main

import "fmt"

func sendJobs(jobs chan<- string) {
	defer close(jobs)
	for i := range 4 {
		jobs <- fmt.Sprintf("job-%v", i)
		fmt.Println("sent job", i)
	}

	fmt.Println("all job sent")
}

func main() {
	jobs := make(chan string, 5)
	done := make(chan bool)

	go func() {
		for {
			if j, more := <-jobs; more {
				fmt.Println("doing:", j)
			} else {
				fmt.Println("no more work")
				done <- true
				break
			}
		}
	}()

	sendJobs(jobs)

	<-done

	j, more := <-jobs
	fmt.Println("more jobs left?", more)
	// Reading from a closed channel succeeds immediately, returning the zero value of the underlying type
	fmt.Println("job got:", j)
	fmt.Println("is j empty?", j == "")
}
