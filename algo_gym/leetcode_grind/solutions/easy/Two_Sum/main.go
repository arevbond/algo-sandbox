package main

func twoSum(nums []int, target int) []int {
	dic := make(map[int]int)
	for i, num := range nums {
		if j, ok := dic[target-num]; ok {
			return []int{i, j}
		}
		dic[num] = i
	}
}
