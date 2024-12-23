package intersectionoftwoarraysii

func intersect(nums1 []int, nums2 []int) []int {
    if len(nums1) < len(nums2) {
        nums1, nums2 = nums2, nums1
    }

    counterNums2 := make(map[int]int)
    for _, n := range nums2 {
        counterNums2[n]++
    }

    result := make([]int, 0)
    for _, n := range nums1 {
        if _, ok := counterNums2[n]; ok {
            result = append(result, n)
            counterNums2[n]--
            if counterNums2[n] == 0 {
                delete(counterNums2, n)
            }
        }
    }
    return result
}
