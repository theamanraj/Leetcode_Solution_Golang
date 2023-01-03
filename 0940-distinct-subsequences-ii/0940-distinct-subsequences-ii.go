func distinctSubseqII(s string) int {
	dp := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = 1 // sub-sequence string(dp[i])
		for j := i - 1; j >= 0; j-- {
			dp[i] += dp[j]
			if s[i] == s[j] {
				// sub-sequence string(dp[i]) already exists, remove duplicate
				dp[i]--
				break
			}
		}
		dp[i] %= 1000000007
	}
	total := 0
	for _, v := range dp {
		total += v
	}
	return total % 1000000007
}