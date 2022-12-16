
func partition(s string) [][]string {
	isPalindrome := func(str string) bool {
		for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
			if str[i] != str[j] {
				return false
			}
		}
		return true
	}
	n := len(s)
	dp := make([][][]string, n)
	// Candidate substring is [begin, end), includes begin, excludes end
	for begin := n - 1; begin >= 0; begin-- {
		for end := begin + 1; end <= n; end++ {
			if isPalindrome(s[begin:end]) {
				if end == n {
					dp[begin] = append(dp[begin], append([]string{}, s[begin:end]))
					continue
				}
				for _, each := range dp[end] {
					newEach := append([]string{}, s[begin:end])
					newEach = append(newEach, each...)
					dp[begin] = append(dp[begin], newEach)
				}
			}
		}
	}
	return dp[0]
}