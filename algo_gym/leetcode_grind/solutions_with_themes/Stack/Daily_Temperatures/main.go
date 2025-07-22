package main

import "fmt"

type Stack struct {
	values [][]int
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(val []int) {
	s.values = append(s.values, val)
}

func (s *Stack) Pop() ([]int, bool) {
	if len(s.values) == 0 {
		return []int{}, false
	}
	val := s.values[len(s.values)-1]
	s.values = s.values[:len(s.values)-1]
	return val, true
}

func dailyTemperatures(temperatures []int) []int {
    answer := make([]int, len(temperatures))
    for i := 0; i < len(answer); i++ {
    	answer[i] = 0
    }

    stack := NewStack()
    for i, temp := range temperatures {
    	for {
    		if len(stack.values) == 0 {
    			stack.Push([]int{temp, i})
    			break
    		}
    		last, _ := stack.Pop()
    		lastValue, lastIndex := last[0], last[1]
    		if temp > lastValue {
    			answer[lastIndex] = i - lastIndex
    		} else {
    			stack.Push([]int{lastValue, lastIndex})
    			stack.Push([]int{temp, i})
    			break
    		}
    	}	
    }

    return answer
}

func main() {
	fmt.Println(dailyTemperatures([]int{73,74,75,71,69,72,76,73}))
	fmt.Println(dailyTemperatures([]int{30, 40, 50, 60}))
	fmt.Println(dailyTemperatures([]int{30, 60, 90}))
}