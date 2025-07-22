package main

import "fmt"

func productExceptSelf(nums []int) []int {
	forwardArray, reverseArray := make([]int, len(nums)), make([]int, len(nums))
	forwardArray[0], reverseArray[len(nums)-1] = 1, 1
	for i := 1; i < len(nums); i++ {
		forwardArray[i] = forwardArray[i-1] * nums[i-1]
	}
	for i := len(nums) - 1; i >= 0; i-- {
		reverseArray[i] = reverseArray[i+1] * nums[i+1]
	}
	answer := make([]int, len(nums))
	for i, _ := range answer {
		answer[i] = forwardArray[i] * reverseArray[i]
	}
	return answer
}

func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
}
