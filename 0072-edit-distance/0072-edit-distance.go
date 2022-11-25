func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func minDistance(word1 string, word2 string) int {
	if len(word2) < len(word1) {
		word1, word2 = word2, word1
	}

	dp := make([]int, len(word2)+1)
	for i := 0; i <= len(word2); i++ {
		dp[i] = i
	}
	for i := 1; i <= len(word1); i++ {
		ndp := make([]int, len(word2)+1)
		ndp[0] = i
		for j := 1; j <= len(word2); j++ {
			c := 1
			if word1[i-1] == word2[j-1] {
				c = 0
			}
			ndp[j] = min(
				dp[j-1]+c,
				min(dp[j], ndp[j-1])+1,
			)
		}
		dp = ndp
	}
	return dp[len(word2)]
}