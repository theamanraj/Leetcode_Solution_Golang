func numWays(steps int, arrLen int) int {
	if arrLen > steps/2+1 {
		arrLen = steps/2 + 1
	}
	dp := make([][]int, steps+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, arrLen)
	}

	dp[0][0] = 1
	for i := 1; i <= steps; i++ {
		for j := 0; j < arrLen; j++ {
			dp[i][j] = (dp[i][j] + dp[i-1][j]) % (1e9 + 7)
			if j-1 >= 0 {
				dp[i][j] = (dp[i][j] + dp[i-1][j-1]) % (1e9 + 7)
			}
			if j+1 < arrLen {
				dp[i][j] = (dp[i][j] + dp[i-1][j+1]) % (1e9 + 7)
			}
		}
	}
	return dp[steps][0]
}