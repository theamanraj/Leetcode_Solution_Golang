func maxProfit(prices []int) int {

	prices = append(prices, -1)  // append -1 is for final loop to add the last valley-peak profit
	profit := 0
	valley := prices[0]

	for i := 1; i < len(prices); i++ {
		if prices[i-1] > prices[i] {
            
			profit += prices[i-1] - valley
			valley = prices[i]
            
		}
	}

	return profit

}