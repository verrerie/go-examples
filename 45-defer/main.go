package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	path := filepath.Join(os.TempDir(), "defer.txt")
	f := createFile(path)
	defer closeFile(f)
	writeFile(f)
	readFile(path)
}

func readFile(path string) {
	fmt.Println("reading")
	bytes, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	fmt.Printf("content: %v\n", bytes)
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintf(f, "this is file %v\n", f.Name())
}

func closeFile(f *os.File) {
	fmt.Println("closing")
	err := f.Close()
	if err != nil {
		panic(err)
	}
}

func createFile(path string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	return f
}
