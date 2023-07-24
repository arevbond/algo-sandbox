package main

import "fmt"

func maxProfit(prices []int) int {
    maximumProfit := 0
    minimumPrice := prices[0]
    for i := 1; i < len(prices); i++ {
    	currentProfit := prices[i] - minimumPrice
    	maximumProfit = max(maximumProfit, currentProfit)
    	minimumPrice = min(minimumPrice, prices[i])
    }
    return maximumProfit
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maxProfit([]int{7,1,5,3,6,4}))
	fmt.Println(maxProfit([]int{7,6,4,3,1}))
}