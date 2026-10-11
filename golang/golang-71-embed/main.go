package main

import (
	"embed"
	"fmt"
)

//go:embed folder/single_file.txt
var fileString string

//go:embed folder/single_file.txt
var fileBytes []byte

//go:embed folder/single_file.txt
//go:embed folder/**.hash
var folder embed.FS

func main() {
	fmt.Print(fileString)
	fmt.Print(string(fileBytes))

	c1, _ := folder.ReadFile("folder/file1.hash")
	fmt.Print(string(c1))

	c2, _ := folder.ReadFile("folder/file2.hash")
	fmt.Print(string(c2))
}
