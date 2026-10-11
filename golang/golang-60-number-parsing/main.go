package main

import (
	"fmt"
	"strconv"
)

func main() {
	f, _ := strconv.ParseFloat("1.2345", 64)
	fmt.Println(f)

	i, _ := strconv.ParseInt("1333", 0, 64)
	fmt.Println(i)

	d, _ := strconv.ParseInt("0x1ac", 0, 64)
	fmt.Println(d)

	u, _ := strconv.ParseUint("0xa", 0, 64)
	fmt.Println(u)

	k, _ := strconv.Atoi("99")
	fmt.Println(k)

	_, err := strconv.Atoi("hello")
	fmt.Println(err)
}
