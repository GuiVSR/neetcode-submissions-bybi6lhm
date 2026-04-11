func maxProfit(prices []int) int {
	left := 0
	right := 1
	maxProfit := 0

	for right < len(prices) {
		transaction := prices[right] - prices[left]

		if transaction > maxProfit {
			maxProfit = transaction
		}

		if prices[right] < prices[left] {
			left = right
		}

		right++
	} 

	return maxProfit
}
