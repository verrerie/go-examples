package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func ck(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	const dir = "subdir"
	e := os.Mkdir(dir, 0o755)
	ck(e)
	defer os.RemoveAll(dir)

	newEmptyFile := func(name string) {
		d := []byte("")
		ck(os.WriteFile(name, d, 0o644))
	}

	newEmptyFile(filepath.Join(dir, "file1"))
	e = os.MkdirAll(filepath.Join(dir, "parent", "child"), 0o755)
	ck(e)

	newEmptyFile(filepath.Join(dir, "parent", "file3"))
	newEmptyFile(filepath.Join(dir, "parent", "file4"))
	newEmptyFile(filepath.Join(dir, "parent", "child", "file5"))

	c, e := os.ReadDir(filepath.Join(dir, "parent"))
	ck(e)

	fmt.Println("listing sub-folders")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	e = os.Chdir(filepath.Join(dir, "parent", "child"))
	ck(e)

	c, e = os.ReadDir(".")
	ck(e)

	fmt.Println("listing parent/child/")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	e = os.Chdir("../../..")
	ck(e)

	fmt.Println("visiting subdir")
	e = filepath.WalkDir(dir, visit)
	ck(e)
}

func visit(path string, d fs.DirEntry, err error) error {
	if err != nil {
		return err
	}
	fmt.Println(" ", path, d.IsDir())
	return nil
}
