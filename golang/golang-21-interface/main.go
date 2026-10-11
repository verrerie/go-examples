package main

import (
	"fmt"
	"math"
)

type geo interface {
	area() int
	perim() int
}

type rect struct {
	w int
	h int
}

func (r *rect) area() int {
	return r.w * r.h
}

func (r *rect) perim() int {
	return 2 * (r.h + r.w)
}

type circle struct {
	r int
}

func (c *circle) area() int {
	return int(math.Pi * float32(c.r) * float32(c.r))
}

func (c *circle) perim() int {
	return int(2.0 * math.Pi * float32(c.r))
}

func measure(g geo) {
	fmt.Println(g)
	fmt.Println("area:", g.area())
	fmt.Println("perim:", g.perim())
}

func identifyCircle(g geo) {
	if c, ok := g.(*circle); ok {
		fmt.Println("circle with r:", c.r)
	} else {
		fmt.Println("not a circle:", g)
	}
}

func main() {
	g1 := rect{5, 6}
	g2 := circle{3}

	measure(&g1)
	measure(&g2)

	identifyCircle(&g1)
	identifyCircle(&g2)
}
