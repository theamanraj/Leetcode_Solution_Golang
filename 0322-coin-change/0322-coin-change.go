func coinChange(coins []int, amount int) int {
	sort.Ints(coins)
	dp := make([]int, amount+1, amount+1)
	uniqueCoins := make([]int, len(coins)+1)
	maxCoins := amount + 1
	for _, coin := range coins {
		if coin <= amount && dp[coin] == 0 {
			dp[coin] = 1
			uniqueCoins = append(uniqueCoins, coin)
		}
	}
	for money := 1; money <= amount; money++ {
		if dp[money] != 0 {
			continue
		}
		dp[money] = -1
		min := maxCoins
		for _, c := range uniqueCoins {
			if c < money {
				a := dp[money-c]
				if a > 0 && a+1 < min {
					min = a + 1
					dp[money] = min
				}
			} else {
				break
			}
		}
	}
	//fmt.Println(dp[1:])
	return dp[amount]
}