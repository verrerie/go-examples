package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	fmt.Print(rand.IntN(100), ",")
	fmt.Println(rand.IntN(100))

	fmt.Print(rand.Float64()*5+5, ",")
	fmt.Print(rand.Float64()*5+5, ",")
	fmt.Println(rand.Float64())

	seed1 := rand.NewPCG(42, 1024)
	r2 := rand.New(seed1)
	fmt.Print(r2.IntN(100), ",")
	fmt.Println(r2.IntN(100))

	seed2 := rand.NewPCG(42, 1024)
	r3 := rand.New(seed2)
	fmt.Print(r3.IntN(100), ",")
	fmt.Println(r3.IntN(100))
}
