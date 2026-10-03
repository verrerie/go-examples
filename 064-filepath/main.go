package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {
	p := filepath.Join("dir1", "dir2", "filename")
	fmt.Println("p:", p)

	fmt.Println(filepath.Join("dir1//", "file1"))
	fmt.Println(filepath.Join("dir1/../dir1", "file2"))

	fmt.Println("dir:", filepath.Dir(p))
	fmt.Println("base:", filepath.Base(p))

	fmt.Println("abs?", filepath.IsAbs("dir1/dir2"))
	fmt.Println("abs?", filepath.IsAbs("/dir1/file2"))

	fmt.Println("ext:", filepath.Ext("test.json"))
	fmt.Println("ext:", filepath.Ext("test"))

	fmt.Println("trimed:", strings.TrimSuffix("setup.xml", filepath.Ext("setup.xml")))

	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)
	rel, err = filepath.Rel("a/b", "a/c/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)
}
