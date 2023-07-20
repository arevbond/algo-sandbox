package main

import "fmt"

func longestConsecutive(nums []int) int {
	set := make(map[int]bool)
	for _, num := range nums {
		set[num] = true
	}
	maxLen := 0

	for _, num := range nums {
		if _, ok := set[num-1]; !ok {
			curLen := 1
			nextNum := num + 1
			for {
				if _, ok2 := set[nextNum]; ok2 {
					curLen++
					nextNum++
				} else {
					break
				}
			}
			maxLen = max(maxLen, curLen)
		}
	}
	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
	fmt.Println(longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}))
}
