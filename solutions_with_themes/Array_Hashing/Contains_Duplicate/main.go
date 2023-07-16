package main

func containsDuplicate(nums []int) bool {
	set := make(map[int]bool)
	for _, num := range nums {
		if _, ok := set[num]; ok {
			return true
		}
		set[num] = true
	}
	return false
}
