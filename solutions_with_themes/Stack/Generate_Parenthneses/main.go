package main

import (
	"fmt"
	"strings"
)

func generateParenthesis(n int) []string {
	stack := make([]string, 0)
	res := make([]string, 0)
	var backtrack func(openN, closedN int)
	backtrack = func(openN, closedN int) {
		if openN == closedN && closedN == n {
			str := strings.Join(stack, "")
			res = append(res, str)
			return
		}
		if openN < n {
			stack = append(stack, "(")
			backtrack(openN+1, closedN)
			stack = stack[:len(stack)-1]
		}
		if closedN < openN {
			stack = append(stack, ")")
			backtrack(openN, closedN+1)
			stack = stack[:len(stack)-1]
		}
	}
	backtrack(0, 0)
	return res
}

func main() {
	//fmt.Println(generateParenthesis(1))
	//fmt.Println(generateParenthesis(2))
	fmt.Println(generateParenthesis(3))
}
