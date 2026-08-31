package main

import (
	"errors"
	"fmt"
)

func f(v int) (int, error) {
	if v == 42 {
		return -1, errors.New("cannot work with 42")
	}
	return v + 3, nil
}

var (
	ErrorOutOfTea = errors.New("not enough tea left")
	ErrorPower    = errors.New("not enough power left")
)

func makeTea(v int) error {
	switch v {
	case 2:
		return ErrorOutOfTea
	case 4:
		return fmt.Errorf("cannot make tea: %w", ErrorPower)
	}
	fmt.Println("making tea:", v)
	return nil
}

func main() {
	v, e := f(42)
	fmt.Println("v, e=", v, e)
	v, e = f(67)
	fmt.Println("v, e=", v, e)

	e = makeTea(2)
	fmt.Println(e)
	e = makeTea(4)

	fmt.Println(e)
	e = makeTea(6)
	fmt.Println(e)

	for i := 1; i < 10; i++ {
		e = makeTea(i)
		if e != nil {
			fmt.Println("couldn't make tea:", e)
		}
	}
}
