package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	rs := Constructor()
	fmt.Println(rs.Insert(1))
	fmt.Println(rs.Insert(2))
	fmt.Println(rs.Insert(3))
	fmt.Println(rs.Insert(4))
	fmt.Println(rs.values)
	fmt.Println(rs.GetRandom())
	fmt.Println(rs.Remove(5))
	fmt.Println(rs.values)
	fmt.Println(rs.Remove(4))
	fmt.Println(rs.values)
	fmt.Println(rs.Remove(3))
	fmt.Println(rs.values)
	fmt.Println(rs.Remove(2))
	fmt.Println(rs.values)
	fmt.Println(rs.GetRandom())
	fmt.Println(rs.Remove(1))
	fmt.Println(rs.values)
	fmt.Println(rs.Remove(1))
	fmt.Println(rs.Insert(10))
	fmt.Println(rs.Insert(20))
	fmt.Println(rs.Insert(30))
	fmt.Println(rs.Insert(40))
	fmt.Println(rs.Remove(20))
	fmt.Println(rs.values)
	fmt.Println(rs.Insert(30))
}

type RandomizedSet struct {
	// space = O(n), n is the current number of elements
	positions map[int]int
	values    []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		positions: make(map[int]int),
		values:    []int{},
	}
}

func (rs *RandomizedSet) Insert(val int) bool {
	// time = amortized O(1) as occasionally involves allocation
	// space = O(1)
	if _, ok := rs.positions[val]; ok {
		return false
	} else {
		p := len(rs.values)
		rs.values = append(rs.values, val)
		rs.positions[val] = p
		return true
	}
}

func (rs *RandomizedSet) Remove(val int) bool {
	// time = O(1)
	// space = O(1)
	if _, ok := rs.positions[val]; ok {
		// last element can be removed without touching the rest
		// but we need to keep its value
		index := rs.positions[val]
		last := len(rs.values) - 1
		lastValue := rs.values[last]
		// first update the to-be removed position to last
		rs.values[index] = lastValue
		rs.positions[lastValue] = index
		// remove the last element
		delete(rs.positions, val)
		rs.values = rs.values[:last]
		return true
	} else {
		return false
	}
}

func (rs *RandomizedSet) GetRandom() int {
	// time = O(1)
	// space = O(1)
	if len(rs.values) == 0 {
		panic("set is empty")
	}
	return rs.values[rand.IntN(len(rs.values))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
