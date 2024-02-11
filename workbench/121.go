package workbench

func MaxProfit(nums []int) int {
	var profit int

	buy := nums[0]
	for _, sell := range nums[1:] {
		if sell > buy {
			profit = max(profit, sell-buy)
		} else {
			buy = sell
		}
	}

	return profit
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
