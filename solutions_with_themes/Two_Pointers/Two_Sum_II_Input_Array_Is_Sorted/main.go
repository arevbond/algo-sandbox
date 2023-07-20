package main

func twoSum(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1
	for l < r {
		curSum := numbers[l] + numbers[r]
		if curSum == target {
			return []int{l + 1, r + 1}
		}
		if curSum > target {
			r--
		} else {
			l++
		}
	}
	return []int{}
}
