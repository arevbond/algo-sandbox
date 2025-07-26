package tasks

func twoSum(nums []int, target int) []int {
	// number: indx
	mp := make(map[int]int)

	for i, num := range nums {
		if prevIndx, ok := mp[target-num]; ok {
			return []int{prevIndx, i}
		}

		mp[num] = i
	}

	return nil
}
