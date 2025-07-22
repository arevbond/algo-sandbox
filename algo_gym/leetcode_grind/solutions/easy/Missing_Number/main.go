package main

func missingNumber(nums []int) int {
	result := 0
	for i, num := range nums {
		result ^= i + 1
		result ^= num
	}
	return result
}
