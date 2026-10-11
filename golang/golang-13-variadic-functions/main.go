package main

import "fmt"

func main() {
	concat := func(ss ...string) string {
		result := ""
		for _, s := range ss {
			result = result + s
		}
		return result
	}
	fmt.Println("concated:", concat("1", "22", "333", "4444", "end"))
}
