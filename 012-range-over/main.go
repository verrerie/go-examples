package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	for i := range nums {
		fmt.Println(i, "th element")
	}

	strs := "hello world!"
	for _, c := range strs {
		fmt.Println("char:", c)
	}
}
