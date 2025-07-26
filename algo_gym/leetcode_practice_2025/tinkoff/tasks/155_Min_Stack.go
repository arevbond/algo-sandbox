package tasks

// Must implement O(1) time complexity for each method
type MinStack struct {
	vals        []int
	orderedVals []int
}

func Constructor() MinStack {
	return MinStack{vals: make([]int, 0), orderedVals: make([]int, 0)}
}

func (s *MinStack) Push(val int) {
	s.vals = append(s.vals, val)

	if len(s.orderedVals) == 0 {
		s.orderedVals = append(s.orderedVals, val)
	} else {
		if s.orderedVals[len(s.orderedVals)-1] >= val {
			s.orderedVals = append(s.orderedVals, val)
		}
	}
}

func (s *MinStack) Pop() {
	if len(s.vals) == 0 {
		return
	}

	last := s.vals[len(s.vals)-1]

	s.vals = s.vals[:len(s.vals)-1]

	if last == s.orderedVals[len(s.orderedVals)-1] {
		s.orderedVals = s.orderedVals[:len(s.orderedVals)-1]
	}
}

// never used in empty stack
func (s *MinStack) Top() int {
	if len(s.vals) == 0 {
		return -1
	}

	return s.vals[len(s.vals)-1]
}

func (s *MinStack) GetMin() int {
	if len(s.orderedVals) == 0 {
		return -1
	}

	return s.orderedVals[len(s.orderedVals)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
