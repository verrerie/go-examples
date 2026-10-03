package main

import "fmt"

type person struct {
	name string
	age  int
}

func newP(name string) *person {
	p := person{name: name}
	p.age = 41
	return &p
}

func main() {
	p1 := newP("Alice")
	fmt.Println(*p1)
	p1.age = 39
	fmt.Println(*p1)

	fmt.Println(person{"Bob", 30})
	fmt.Println(person{name: "Cal", age: 32})
	fmt.Println(newP("Dan"))

	dog := struct {
		name   string
		isGood bool
	}{
		"Alex",
		true,
	}
	fmt.Println(dog)
}
