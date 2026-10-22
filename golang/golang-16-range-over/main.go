package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	for i := range nums {
		fmt.Println(i, "th element")
	}

	strs := "hello world!"
	fmt.Println("runes")
	for _, c := range strs {
		fmt.Println("char:", c)
	}

	fmt.Println("types")
	for _, b := range []byte(strs) {
		fmt.Println("byte:", b)
	}
}
