package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var op atomic.Uint64

	var wg sync.WaitGroup

	for range 1000 {
		wg.Go(func() {
			for range 10 {
				op.Add(1)
			}
		})
	}

	wg.Wait()

	fmt.Println("result:", op.Load())
}
