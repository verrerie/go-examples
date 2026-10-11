package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	s := "sha256 this tring"

	h := sha256.New()

	h.Write([]byte(s))

	bs := h.Sum(nil)

	fmt.Printf("original: %v\n  hashed: %x\n", s, bs)
}
