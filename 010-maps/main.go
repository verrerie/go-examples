package main

import "fmt"

func main() {
	m := make(map[string]int)
	fmt.Println("m's len:", len(m))

	m["a"] = 1
	m["b"] = 2
	m["c"] = 100
	fmt.Println("m:", m)
	fmt.Println("a", m["a"])
	fmt.Println("b", m["b"])
	fmt.Println("c", m["c"])

	delete(m, "a")
	delete(m, "c")
	fmt.Println("m:", m)

	m2 := map[string]int{"k1": 100, "k2": 200, "k3": 500}
	v, exists := m2["k2"]
	fmt.Println("v, exits=", v, exists)
	for k, v := range m2 {
		fmt.Println("k, v=", k, v)
	}
}
