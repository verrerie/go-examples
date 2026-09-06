package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	os.Setenv("FOO", "BAR")
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))

	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Println(pair[0])
	}
	fmt.Println("total:", len(os.Environ()))
}
