func moveZeroes(nums []int) {
	zIndx := 0
	for r := 0; r < len(nums); r++ {
		if zIndx >= len(nums) {
			return
		}

		if r > zIndx && nums[r] != nums[zIndx] && nums[zIndx] == 0 {
			nums[r], nums[zIndx] = nums[zIndx], nums[r]
		}

		for ; zIndx < len(nums) && nums[zIndx] != 0; zIndx++ {}
	}
}
