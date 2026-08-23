package main

import "fmt"

func main() {
	var a [4]int
	fmt.Println("emp:", a)

	a[3] = 100
	fmt.Println("a:", a)
	fmt.Println("a[3]:", a[3])
	fmt.Println("len:", len(a))

	b := [4]int{1, 2, 3, 4}
	fmt.Println("b:", b)
	fmt.Println("b[3]:", b[3])
	fmt.Println("len(b):", len(b))

	c := [...]int{1, 3, 5, 7}
	fmt.Println("c:", c)

	d := [...]int{1, 3: 100, 200, 500}
	fmt.Println("d:", d)

	twoD := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println("twoD:", twoD)
}
