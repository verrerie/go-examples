package main

import (
	"fmt"
	"strconv"
)

type wheel struct {
	wheelType string
	num       int
}

func (w *wheel) describe() string {
	return strconv.Itoa(w.num) + " " + w.wheelType + " wheels"
}

type engine struct {
	engineType string
	power      int
}

func (e *engine) String() string {
	return "engine of type " + e.engineType + ", power=" + strconv.Itoa(e.power)
}

type car struct {
	*wheel
	*engine
	price int
	year  int
}

func main() {
	c1 := car{
		&wheel{
			"racing",
			4,
		},
		&engine{
			"electric",
			1900,
		},
		25000,
		2026,
	}

	fmt.Println("car 1, price =", c1.price, "year =", c1.year)
	fmt.Println(c1.engine)
	fmt.Println(c1.describe())

	type desc interface {
		describe() string
	}

	var d desc = c1
	fmt.Println("this is", d.describe())
}
