package main

import (
	"fmt"
)

func findMin(nums []int) int {
    l, r := 0, len(nums) - 1
    for l < r {
    	m := (l + r) / 2
    	if nums[m] > nums[r] {
    		l = m + 1
    	} else {
    		r = m
    	}
    }
    return nums[l]
}

func main() {
	fmt.Println(findMin([]int{3,4,5,1,2}))
	fmt.Println(findMin([]int{4,5,6,7,0,1,2}))
	fmt.Println(findMin([]int{11,13,15,17}))
	fmt.Println(findMin([]int{2, 1}))
}