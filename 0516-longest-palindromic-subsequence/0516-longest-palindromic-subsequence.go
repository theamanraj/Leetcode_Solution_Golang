func longestPalindromeSubseq(s string) int {
    dp := make([]int, len(s))
	for i := range dp {
		dp[i] = 1
	}
	for i := len(s) - 2; i >= 0; i-- {
		pre := 0
		for j := i + 1; j < len(s); j++ {
			temp := dp[j]
			if s[i] == s[j] {
				dp[j] = pre + 2
			} else {
				dp[j] = int(math.Max(float64(dp[j-1]), float64(dp[j])))
			}
			pre = temp
		}
	}
	return dp[len(s)-1]
}