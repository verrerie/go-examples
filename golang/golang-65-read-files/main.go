package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func ck(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	path := filepath.Join(".", "data")
	data, e := os.ReadFile(path)
	ck(e)
	fmt.Println(string(data))

	f, e := os.Open(path)
	ck(e)
	defer f.Close()

	b1 := make([]byte, 6)
	n, e := f.Read(b1)
	ck(e)
	fmt.Println("read", n, "bytes")
	fmt.Println(string(b1[:n]))

	o2, e := f.Seek(3, io.SeekStart)
	ck(e)
	b2 := make([]byte, 2)
	n2, e := f.Read(b2)
	ck(e)
	fmt.Println("read", n2, "bytes at", o2)
	fmt.Println(string(b2[:n2]))

	_, e = f.Seek(2, io.SeekCurrent)
	ck(e)
	_, e = f.Seek(-4, io.SeekEnd)
	ck(e)
	o3, e := f.Seek(3, io.SeekStart)
	ck(e)

	b3 := make([]byte, 3)
	n3, e := io.ReadAtLeast(f, b3, 2)
	ck(e)
	fmt.Printf("read %v bytes at %d: %v\n", n3, o3, string(b3))

	_, e = f.Seek(0, io.SeekStart)
	ck(e)

	r4 := bufio.NewReader(f)
	b4, e := r4.Peek(5)
	ck(e)
	fmt.Printf("read %d bytes: %v", len(b4), string(b4))
}
