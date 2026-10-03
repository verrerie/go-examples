package main

import "fmt"

const (
	a1 = iota
	a2
	a3
)

const (
	b1, c1 = 1 << iota, 3 * iota
	b2, c2
	b3, c3
)

type ServerState int

const (
	StateIdle ServerState = iota
	StateConnected
	StateError
	StateRetrying
)

var state = map[ServerState]string{
	StateIdle:      "idle",
	StateConnected: "connected",
	StateError:     "error",
	StateRetrying:  "retrying",
}

func (ss ServerState) String() string {
	return state[ss]
}

func next(s ServerState) ServerState {
	switch s {
	case StateIdle:
		return StateConnected
	case StateConnected, StateRetrying:
		return StateIdle
	case StateError:
		return StateRetrying
	default:
		panic(fmt.Errorf("unknown state: %s", s))
	}
}

func main() {
	fmt.Println("a1, a2, a3=", a1, a2, a3)
	fmt.Println("b1, b2, b3=", b1, b2, b3)
	fmt.Println("c1, c2, c3=", c1, c2, c3)

	fmt.Println(StateIdle, "->", next(StateIdle))
	fmt.Println(StateConnected, "->", next(StateConnected))
	fmt.Println(StateError, "->", next(StateError))
	fmt.Println(StateRetrying, "->", next(StateRetrying))
}
