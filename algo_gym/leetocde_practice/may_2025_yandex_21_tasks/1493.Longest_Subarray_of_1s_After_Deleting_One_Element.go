func longestSubarray(nums []int) int {
	prevOnes, curOnes, result := 0, 0, 0

	for _, num := range nums {
		if num == 1 {
			curOnes++
			continue
		}

		result = max(result, prevOnes + curOnes)
		prevOnes, curOnes = curOnes, 0
	}

	if curOnes == len(nums) {
		curOnes--
	}

	return max(result, prevOnes+curOnes)
}
