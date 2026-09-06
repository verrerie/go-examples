package main

import (
	"fmt"
	"os"
)

var (
	p  = fmt.Printf
	pl = fmt.Println
)

type Point struct {
	x int
	y int
}

func main() {
	p1 := Point{10, 100}
	p("point: %v", p1)
	pl()
	p("point: %+v", p1)
	pl()
	p("point: %#v", p1)
	pl()
	p("point: %T", p1)
	pl()
	p("boolean: %t", true)
	pl()
	p("int: %d", 10)
	pl()
	p("bin: %b", 8)
	pl()
	p("str: %s", "hello")
	pl()
	p("float: %.2f", 12.3)
	pl()
	p("float: %e", 12.3)
	pl()
	p("float: %E", 12.3)
	pl()
	p("str2: %q", "\"string\"")
	pl()
	p("str2: %s", "\"string\"")
	pl()
	p("hex: %x", 10)
	pl()
	p("hex this: %x", "abc")
	pl()
	p("ponter: %p", &p)
	pl()
	p("|%9d|%9.2f|", 12, 1.2)
	pl()
	p("|%9s|%9s|", "hello", "world")
	pl()
	p("|%-9s|%-9s|", "great", "news")
	pl()
	p("hello: %s", fmt.Sprintf("Nick %d", 120))
	pl()
	fmt.Fprintf(os.Stderr, "hey!")
	pl()
}
