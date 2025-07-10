func maxProfit(prices []int) int {
	totalProfit := 0


	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			totalProfit += prices[i] - prices[i-1]
		}
	}

	return totalProfit
}
