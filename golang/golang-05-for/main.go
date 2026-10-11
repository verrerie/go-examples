package main

import "fmt"

func main() {
	i := 0
	for i < 3 {
		fmt.Println(i)
		i = i + 1
	}

	for i := 1; i < 3; i++ {
		fmt.Println(i)
	}

	for i := range 6 {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}

	for {
		fmt.Println("loop")
		break
	}
}
