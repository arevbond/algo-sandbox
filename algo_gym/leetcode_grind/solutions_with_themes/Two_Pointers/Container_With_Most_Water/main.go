package main

import "fmt"

func maxArea(height []int) int {
	l, r := 0, len(height)-1
	maxAmount := 0
	for l < r {
		curAmount := min(height[l], height[r]) * (r - l)
		maxAmount = max(curAmount, maxAmount)
		if height[l] > height[r] {
			r--
		} else {
			l++
		}
	}
	return maxAmount
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}
