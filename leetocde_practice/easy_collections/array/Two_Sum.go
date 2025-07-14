func twoSum(nums []int, target int) []int {
	// value - indx
	mp := make(map[int]int)

	for i, num := range nums {
		if i2, ok := mp[target-num]; ok {
			return []int{i, i2}
		}

		mp[num] = i
	}

	return nil
}
