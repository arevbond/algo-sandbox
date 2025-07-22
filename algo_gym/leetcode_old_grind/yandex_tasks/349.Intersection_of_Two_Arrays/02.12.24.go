package intersectionoftwoarrays

/*

    nums1 = [1,2,2,1], nums2 = [2,2]
    1. set := {}
    2. for n range min(nums1, nums2)
    3. set.add(n)
    4. for n range max(nums1, nums2)
    5. if n exists in set => append n in result

*/

// space: O(K); k = min(len(nums1), len(nums2))
// time: O(N); n = max(len(nums1), len(nums2))
func intersection(nums1 []int, nums2 []int) []int {
    if len(nums1) > len(nums2) {
        nums1, nums2 = nums2, nums1
    }

    set := make(map[int]struct{})
    for _, n := range nums1 {
        set[n] = struct{}{}
    }

    result := make([]int, 0, len(set))
    for _, n := range nums2 {
        if _, ok := set[n]; ok {
            result = append(result, n)
            delete(set, n)
        }
    }
    return result
}
