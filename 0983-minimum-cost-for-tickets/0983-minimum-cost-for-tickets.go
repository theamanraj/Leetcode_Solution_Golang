func mincostTickets(days []int, costs []int) int {
	if len(days) == 0 {
		return 0
	}

	dict := make(map[int]bool)
	for i := 0; i < len(days); i++ {
		dict[days[i]] = true
	}
    dp := make([]int, days[len(days) - 1]+1)
	for i := 1; i < len(dp); i++ {
		if !dict[i] {
			dp[i] = dp[i-1]
			continue
		}

		mx := dp[i-1] + costs[0]
		if i >= 7 {
			mx = min(mx, costs[1]+dp[i-7])
		} else {
			mx = min(mx, costs[1])
		}
		if i >= 30 {
			mx = min(mx, costs[2]+dp[i-30])
		} else {
			mx = min(mx, costs[2])
		}
		dp[i] = mx
	}
    return dp[len(dp) - 1]
}
func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}