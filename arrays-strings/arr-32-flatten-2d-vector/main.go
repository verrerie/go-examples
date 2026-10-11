package main

import "fmt"

func main() {
	v2d := Constructor([][]int{
		{1},
		{},
		{2, 3, 4},
		{5},
	})
	for v2d.HasNext() {
		fmt.Println(v2d.Next())
	}
	v2d = Constructor([][]int{})
	fmt.Println(v2d.HasNext())
}

type Vector2D struct {
	vec [][]int
	r   int
	c   int
}

func Constructor(vec [][]int) Vector2D {
	return Vector2D{vec, 0, 0}
}

func (v *Vector2D) Next() int {
	// time = O(R)
	// space = O(1)
	if v.HasNext() {
		val := v.vec[v.r][v.c]
		v.c++
		return val
	}
	panic("no more element")
}

func (v *Vector2D) HasNext() bool {
	// time = O(R)
	// space = O(1)
	for v.r < len(v.vec) && v.c == len(v.vec[v.r]) {
		v.r++
		v.c = 0
	}
	return v.r != len(v.vec)
}

/**
 * Your Vector2D object will be instantiated and called as such:
 * obj := Constructor(vec);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
