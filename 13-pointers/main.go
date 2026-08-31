package main

import "fmt"

func main() {
	zeroval := func(val int) {
		val = 0
	}
	zeroptr := func(ptr *int) {
		*ptr = 0
	}

	i := 1
	fmt.Println("initial=", i)

	zeroval(i)
	fmt.Println("zeroval->", i)

	zeroptr(&i)
	fmt.Println("zeroptr->", i)

	fmt.Println("pointer=", &i)

	p := new(42)
	fmt.Println("value at p:", *p)
	zeroptr(p)
	fmt.Println("value at p:", *p)
}
