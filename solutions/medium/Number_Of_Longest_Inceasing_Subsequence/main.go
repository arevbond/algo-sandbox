package main

import "fmt"

func findNumberOfLIS(nums []int) int {
	dp := make(map[int][]int) // key = index, value = [length of LIS, count]
	lenLIS, res := 0, 0 // length of LIS, count of LIS

	for i := len(nums) - 1; i >= 0; i-- {
		maxLen, maxCnt := 1, 1
		for j := i + 1; j < len(nums); j++ {
			if nums[i] < nums[j] {
				pair, _ := dp[j]
				length, count := pair[0], pair[1]
				if length + 1 > maxLen {
					maxLen = length + 1
					maxCnt = count
				} else if length + 1 == maxLen {
					maxCnt += count
				}
			}
		}
		if maxLen > lenLIS {
			lenLIS, res = maxLen, maxCnt
		} else if maxLen == lenLIS {
			res += maxCnt
		}
		dp[i] = []int{maxLen, maxCnt}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(findNumberOfLIS([]int{1,3,5,4,7}))	
	fmt.Println(findNumberOfLIS([]int{2,2,2,2,2}))	
}	
