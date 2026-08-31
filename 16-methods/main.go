package main

import "fmt"

type rect struct {
	width, height int
}

func newRect(w, h int) *rect {
	return &rect{w, h}
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r *rect) perim() int {
	return r.width*2 + r.height*2
}

func main() {
	r := newRect(4, 6)
	fmt.Println("rect:", *r)
	fmt.Println("area:", r.area())
	fmt.Println("perimeter:", r.perim())

	r2 := r
	fmt.Println("area:", r2.area())
	fmt.Println("perimeter:", r2.perim())
}
