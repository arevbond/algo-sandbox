package main

import "fmt"

func lengthOfLIS(nums []int) int {
    LIS := make([]int, len(nums))
    for i := 0; i < len(nums); i++ {
    	LIS[i] = 1
    }

    for i := len(nums) - 1; i >= 0; i-- {
    	for j := i + 1; j < len(nums); j++{
    		if nums[i] < nums[j] {
    			LIS[i] = max(LIS[i], 1 + LIS[j])
    		}
    	}
    }
    maxLen := 0
    for _, val := range LIS {
    	maxLen = max(maxLen, val)
    }
    return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// fmt.Println(lengthOfLIS([]int{10,9,2,5,3,7,101,18}))
	fmt.Println(lengthOfLIS([]int{0,1,0,3,2,3}))
}