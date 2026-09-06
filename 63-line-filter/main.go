package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scan := bufio.NewScanner(os.Stdin)

	for scan.Scan() {
		filtered := strings.ToUpper(scan.Text())
		fmt.Println(filtered)
	}

	if e := scan.Err(); e != nil {
		fmt.Fprintln(os.Stderr, "error:", e)
		os.Exit(1)
	}
}
