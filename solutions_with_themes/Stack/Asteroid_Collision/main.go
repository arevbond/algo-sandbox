package main

import "fmt"

type Stack struct {
	values []int
}

func NewStack() *Stack {
	return &Stack{values: []int{}}
}

func (s *Stack) Push(val int) {
	s.values = append(s.values, val)
}

func (s *Stack) Pop() (int, bool) {
	if len(s.values) > 0 {
		val := s.values[len(s.values)-1]
		s.values = s.values[:len(s.values)-1]
		return val, true
	}
	return 0, false
}

func (s *Stack) Len() int {
	return len(s.values)
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func asteroidCollision(asteroids []int) []int {
	stack := NewStack()
	for _, aster := range asteroids {
		if stack.Len() == 0 {
			stack.Push(aster)
		} else {
			for {
				if stack.Len() == 0 {
					stack.Push(aster)
					break
				}
				prev, _ := stack.Pop()
				if (aster > 0 && prev < 0) || (aster > 0 && prev > 0) ||
					(aster < 0 && prev < 0) {
					stack.Push(prev)
					stack.Push(aster)
					break
				}
				if abs(prev) == abs(aster) {
					break
				} else if abs(prev) > abs(aster) {
					stack.Push(prev)
					break
				} else if abs(prev) < abs(aster) {
					continue
				}
			}
		}
	}
	return stack.values
}

func main() {
	fmt.Println(asteroidCollision([]int{5, 10, -5}))
	fmt.Println(asteroidCollision([]int{8, -8}))
	fmt.Println(asteroidCollision([]int{10, 2, -5}))
	fmt.Println(asteroidCollision([]int{5, 4, -10, -5, 2, -1, 3, 4}))
}
