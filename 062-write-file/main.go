package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func ck(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	d1 := []byte("this is sunday")
	path := filepath.Join(".", "data1")
	e := os.WriteFile(path, d1, 0o644)
	ck(e)

	path2 := filepath.Join(".", "data2")
	f, e := os.Create(path2)
	ck(e)
	defer f.Close()

	d2 := []byte("it's Saturday today!\n")
	n2, e := f.Write(d2)
	ck(e)
	fmt.Println("Wrote", n2, "bytes")

	n3, e := f.WriteString("greatings!")
	ck(e)
	fmt.Println("Wrote", n3, "bytes")

	f.Sync()

	w := bufio.NewWriter(f)
	n4, e := w.WriteString("buffered!")
	ck(e)
	fmt.Println("Wrote", n4, "bytes")

	w.Flush()
}
