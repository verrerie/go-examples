package main

import (
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
	f, err := os.CreateTemp("", "sample")
	ck(err)

	fmt.Println("tmp filename:", f.Name())
	defer os.Remove(f.Name())

	_, err = f.Write([]byte{1, 2, 3, 4})
	ck(err)

	dname, err := os.MkdirTemp("", "tempdir")
	ck(err)
	fmt.Println("tempdir name:", dname)

	defer os.RemoveAll(dname)

	fname := filepath.Join(dname, "file1")
	err = os.WriteFile(fname, []byte("hello world!"), 0o644)
	ck(err)

	out, err := os.ReadFile(fname)
	ck(err)
	fmt.Println("read back:", string(out))
}
