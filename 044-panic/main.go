package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	path := filepath.Join(os.TempDir(), "file")
	f, err := os.Create(path)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("file created:", f.Fd())
	}
	panic("a problem")
}
