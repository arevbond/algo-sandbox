func evalRPN(tokens []string) int {
	stack := make([]int, 0)

	for _, token := range tokens {
		number, err := strconv.Atoi(token)
		if nil == err {
			stack = append(stack, number)
			continue
		}

		firstNum := stack[len(stack) - 2]
		secondNum := stack[len(stack) - 1]
		stack = stack[:len(stack) - 2]

		stack = append(stack, processOperation(token, firstNum, secondNum))
	}

	return stack[0]
}

func processOperation(token string, firstNum, secondNum int) int {
	switch token {
	case "+":
		return firstNum + secondNum
	case "-":	
		return firstNum - secondNum
	case "*":
		return firstNum * secondNum
	case "/":
		return firstNum / secondNum
	}
	return -1
}
