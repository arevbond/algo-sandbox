// Time: O(N), Space: O(N)
func rotate1(nums []int, k int) {
	k = k % len(nums)
	indx := len(nums) - k
	left, right := nums[:indx], nums[indx:]
	result := make([]int, 0, len(left)+len(right))
	result = append(result, right...)
	result = append(result, left...)
	for i := 0; i < len(nums); i++ {
		nums[i] = result[i]
	}
	return
}

func rotate(nums []int, k int) {
	k = k % len(nums)
	reverse(nums)
	indx := k
	reverse(nums[:indx])
	reverse(nums[indx:])
}

func reverse(nums []int) {
	l, r := 0, len(nums) - 1
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}
