package best_time_to_buy_and_sell_stock

// O(n²) - brute force, check every pair of buy and sell days
func maxProfit_1(prices []int) int {
	profit := 0
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			if prices[j]-prices[i] > profit {
				profit = prices[j] - prices[i]
			}
		}
	}
	return profit
}

// O(n) - track minimum price so far, update max profit at each step
func maxProfit(prices []int) int {
	profit := 0
	minPrice := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		}
		if prices[i]-minPrice > profit {
			profit = prices[i] - minPrice
		}
	}
	return profit
}
