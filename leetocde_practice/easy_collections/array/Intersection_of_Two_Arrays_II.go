func intersect(nums1 []int, nums2 []int) []int {
	mp1 := make(map[int]int)
	for _, num := range nums1 {
		mp1[num]++
	}

	mp2 := make(map[int]int)
	for _, num := range nums2 {
		mp2[num]++
	}

	res := make(map[int]int)
	for num, cnt1 := range mp1 {
		if cnt2, ok := mp2[num]; ok {
			res[num] = min(cnt1, cnt2)
		}
	}

	result := make([]int, 0, len(res))
	for num, cnt := range res {
		for cnt > 0 {
			result = append(result, num)
			cnt--
		}
	}
	return result
}
