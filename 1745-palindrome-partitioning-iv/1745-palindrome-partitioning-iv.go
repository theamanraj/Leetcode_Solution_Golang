func checkPartitioning(s string) bool {
	n := len(s)
	dp := make([][]bool, n)
	for i, _ := range dp {
		dp[i] = make([]bool, n)
	}
	for j := 0; j < n; j++ {
		for i := 0; i <= j; i++ {
			if s[i] == s[j] {
				if i + 1 <= j - 1 {
					dp[i][j] = dp[i+1][j-1]
				} else {
					dp[i][j] = true
				}
			} else {
				dp[i][j] = false
			}
		}
	}
	for i := 1; i < n - 1; i++ {
		for j := i; j < n - 1; j++ {
			if dp[0][i-1] && dp[i][j] && dp[j+1][n-1] {
				return true
			}
		}
	}
	return false
}
