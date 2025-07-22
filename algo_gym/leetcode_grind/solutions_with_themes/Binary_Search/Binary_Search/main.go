package main

import (
	"fmt"
)

func search(nums []int, target int) int {
    l, r := 0, len(nums) - 1
    for l < r {
    	m := (l + r) / 2
    	if nums[m] >= target {
    		r = m
    	} else {
    		l = m + 1
    	}
    }

    if nums[l] == target{
    	return l
    }
    return -1
}

func main() {
	fmt.Println(search([]int{-1,0,3,5,9,12}, 9))
	fmt.Println(search([]int{-1,0,3,5,9,12}, 2))
}