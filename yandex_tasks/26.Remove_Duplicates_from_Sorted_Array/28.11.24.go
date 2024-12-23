package removeduplicatesfromsortedarray

// [1,2,3,4,3,1]
//l:        ^           
//r:          ^
func removeDuplicates(nums []int) int {
    l := 1
    for r := 1; r < len(nums); {
        if nums[r] != nums[l-1] {
            nums[r], nums[l] = nums[l], nums[r]
            l++
            r++
            for ; r < len(nums) && nums[r] == nums[r-1]; r++ {}
        } else {
            r++
        }
    }
    return l
}
