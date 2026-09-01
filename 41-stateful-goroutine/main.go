package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"sync/atomic"
	"time"
)

type readOp struct {
	key int
	res chan int
}

type writeOp struct {
	key int
	val int
	res chan bool
}

func main() {
	var readCount uint64
	var writeCount uint64

	reads := make(chan readOp)
	writes := make(chan writeOp)

	state := make(map[int]int)
	go func() {
		for {
			select {
			case read := <-reads:
				read.res <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.res <- true
			}
		}
	}()

	var wg sync.WaitGroup

	for range 101 {
		wg.Go(func() {
			read := readOp{
				key: rand.IntN(5),
				res: make(chan int),
			}
			reads <- read
			<-read.res
			atomic.AddUint64(&readCount, 1)
			time.Sleep(time.Millisecond)
		})
	}

	for range 35 {
		wg.Go(func() {
			write := writeOp{
				key: rand.IntN(5),
				val: rand.IntN(999),
				res: make(chan bool),
			}
			writes <- write
			<-write.res
			atomic.AddUint64(&writeCount, 1)
			time.Sleep(time.Millisecond)
		})
	}

	wg.Wait()
	fmt.Println("read count:", readCount)
	fmt.Println("write count:", writeCount)
	fmt.Println("result:", state)
}
