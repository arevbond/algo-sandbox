package findminimuminrotatedsortedarray

/*
             r 
     l
    [3,4,5,6,7,1,2]
           m


*/

func findMin(nums []int) int {
    l, r := 0, len(nums) - 1
    for l < r {
        m := (l + r) / 2
        if nums[m] > r {
            l = m + 1
        } else {
            r = m
        }
    }
    return nums[l]
}
