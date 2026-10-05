package main

import (
	"fmt"
	"strconv"
)

func main() {
	tokens := []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
	fmt.Println(evalRPN(tokens))
}

// time = O(n)
// space = O(n)
func evalRPN(tokens []string) int {
	stack := make([]int, 0)

	push := func(n int) {
		stack = append(stack, n)
	}

	pop := func() int {
		lastIndex := len(stack) - 1
		last := stack[lastIndex]
		stack = stack[:lastIndex]
		return last
	}

	for _, token := range tokens {
		switch token {
		case "-", "+", "/", "*":
			right := pop()
			left := pop()
			result := binaryOp(token, left, right)
			push(result)
		default:
			push(toI(token))
		}
	}

	return pop()
}

func binaryOp(token string, left, right int) int {
	switch token {
	case "+":
		return left + right
	case "-":
		return left - right
	case "*":
		return left * right
	case "/":
		return left / right
	default:
		panic(fmt.Sprintln("unknown operator", token))
	}
}

func toI(token string) int {
	n, err := strconv.Atoi(token)
	if err != nil {
		panic(err)
	}
	return n
}
