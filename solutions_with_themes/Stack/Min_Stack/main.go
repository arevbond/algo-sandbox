package main

type MinStack struct {
	values    []int
	minValues []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (self *MinStack) Push(val int) {
	self.values = append(self.values, val)

	if len(self.minValues) > 0 && self.minValues[len(self.minValues)-1] > val || len(self.minValues) == 0 {
		self.minValues = append(self.minValues, val)
	}

}

func (self *MinStack) Pop() {
	last := self.values[len(self.values)-1]
	if len(self.minValues) > 0 && last == self.minValues[len(self.minValues)-1] {
		self.minValues = self.minValues[:len(self.minValues)-1]
	}
	self.values = self.values[:len(self.values)-1]
}

func (self *MinStack) Top() int {
	return self.values[len(self.values)-1]
}

func (self *MinStack) GetMin() int {
	if len(self.minValues) != 0 {
		return self.minValues[len(self.minValues)-1]
	}
	return -1
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
