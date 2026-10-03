package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println("slice:", []int{1, 2, 3})

	var a []int
	fmt.Println("a:", a)
	fmt.Println("len:", len(a), "cap:", cap(a))
	fmt.Println(a == nil)

	b := make([]int, 3)
	fmt.Println("b:", b)
	fmt.Println(b == nil)
	b[0] = 1
	b[2] = 0
	b[1] = 2
	fmt.Println("b:", b)

	fmt.Println("len:", len(b), "cap:", cap(b))
	b = append(b, 3)
	fmt.Println("len:", len(b), "cap:", cap(b))
	b = append(b, 4)
	fmt.Println("len:", len(b), "cap:", cap(b))

	c := make([]int, len(b))
	copy(c, b)
	fmt.Println("c:", c)
	fmt.Println("c[1:3]:", c[1:3])
	fmt.Println("c[:4]:", c[:4])
	fmt.Println("c[2:]:", c[2:])

	d := []string{"a", "b", "c"}
	fmt.Println("len:", len(d), "cap:", cap(d))
	fmt.Println("d:", d)

	d2 := []string{"1", "2", "3"}
	copy(d2, d)
	fmt.Println("len:", len(d2), "cap:", cap(d2))
	fmt.Println("d2:", d2)
	fmt.Println("is d equal to d2?", slices.Equal(d, d2))

	threeD := [][][]int{}
	const dim = 6
	for i := range dim {
		inner2 := make([][]int, 0)
		for j := range dim {
			inner1 := make([]int, 0)
			for k := range dim {
				inner1 = append(inner1, i+j+k)
			}
			inner2 = append(inner2, inner1)
		}
		threeD = append(threeD, inner2)
	}
	fmt.Println("threeD:", threeD)
	fmt.Println("threeD len:", len(threeD))
	fmt.Println("twoD len:", len(threeD[0]))
	fmt.Println("oneD len:", len(threeD[0][0]))
}
