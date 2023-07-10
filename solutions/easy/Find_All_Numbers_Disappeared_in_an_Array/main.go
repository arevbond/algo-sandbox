package main

func findDisappearedNumbers(nums []int) []int {
	res := make([]int, len(nums))

	// using result array as sorted array for seen values
	for _, v := range nums {
		res[v-1] = v
	}
	// here we have res array where missing records will be with value = 0
	// scan all values with 0 and "shift" them to head of res array
	j := 0
	for i := 0; i < len(nums); i++ {
		if res[i] == 0 {
			res[j] = i + 1
			j++
		}
	}
	return res[:j]
}
