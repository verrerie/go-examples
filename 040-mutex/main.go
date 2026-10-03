package main

import (
	"fmt"
	"sync"
)

type container struct {
	mu   sync.Mutex
	vals map[string]int
}

func (c *container) inc(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.vals[name] = c.vals[name] + 1
}

func main() {
	con := container{
		vals: map[string]int{"a": 0, "b": 0},
	}

	var wg sync.WaitGroup

	increment := func(name string, n int) {
		for range n {
			wg.Go(func() {
				con.inc(name)
			})
		}
	}

	increment("b", 1000)
	increment("a", 1000)
	increment("a", 1000)
	increment("b", 1000)

	wg.Wait()
	fmt.Println(con.vals)
}
