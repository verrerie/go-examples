package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	worker := func(id int) {
		fmt.Println("starting", id)
		time.Sleep(10 * time.Millisecond)
		fmt.Println("done", id)
	}

	var wg sync.WaitGroup
	for i := range 6 {
		wg.Go(func() {
			worker(i)
		})
	}

	wg.Wait()
}
