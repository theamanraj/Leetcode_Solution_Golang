func minCut(s string) int {
	sInByteSlice := []byte(s)
	dp := make([]int, len(s))
	dp[0] = 0
	// endIndex, inclusive
	for endIndex := 1; endIndex < len(s); endIndex++ {
		if isPalindrome(sInByteSlice[:endIndex+1]) {
			dp[endIndex] = 0
			continue
		}
		dp[endIndex] = dp[endIndex-1] + 1
		for startIndex := endIndex-1; startIndex > 0; startIndex-- {
			if isPalindrome(sInByteSlice[startIndex:endIndex+1]) {
				if tmpCut := dp[startIndex-1] + 1; tmpCut < dp[endIndex] {
					dp[endIndex] = tmpCut
				}
			}
		}
	}
	return dp[len(s)-1]
}

func isPalindrome(s []byte) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}