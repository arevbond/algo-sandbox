package main

import (
	"fmt"
	"strconv"
)

func operation(operator string, x, y int) int {
	if operator == "+" {
		return x + y
	}
	if operator == "-" {
		return x - y
	}
	if operator == "*" {
		return x * y
	}
	if operator == "/" {
		return x / y
	}
	return 0
}

func evalRPN(tokens []string) int {
	stack := []int{}
	for _, elem := range tokens {
		if num, err := strconv.Atoi(elem); err == nil {
			stack = append(stack, num)
		} else {
			firstNum, secondNum := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			newNum := operation(elem, firstNum, secondNum)
			stack = append(stack, newNum)
		}
	}
	return stack[len(stack)-1]
}

func main() {
	fmt.Println(evalRPN([]string{"2", "1", "+", "3", "*"}))
	fmt.Println(evalRPN([]string{"4", "13", "5", "/", "+"}))
	fmt.Println(evalRPN([]string{"10", "6", "9", "3", "+", "-11",
		"*", "/", "*", "17", "+", "5", "+"}))
}
