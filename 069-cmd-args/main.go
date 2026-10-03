package main

import (
	"fmt"
	"os"
)

func main() {
	argsFull := os.Args
	argsWithoutProg := argsFull[1:]

	secondArg := argsFull[2]

	fmt.Println(argsFull)
	fmt.Println(argsWithoutProg)
	fmt.Println(secondArg)
}
