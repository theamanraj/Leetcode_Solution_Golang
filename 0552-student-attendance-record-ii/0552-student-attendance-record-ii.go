func checkRecord(n int) int {
	switch n {
	case 1:
		return 3
	case 2:
		return 8
	case 3:
		return 19
	}
	// dp[i]: count without 'A'
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 2
	dp[2] = 4
	dp[3] = 7
	for i := 4; i <= n; i++ {
        // XXXP: dp[i-1]
        // XXPL: dp[i-2]
        // XPLL: dp[i-3]
		dp[i] = (dp[i-1] + dp[i-2] + dp[i-3]) % 1000000007
	}
	withA := 0
	for i := 0; i <= n-1; i++ {
        // AXXX: dp[0] * dp[3]
        // XAXX: dp[1] * dp[2]
        // XXAX: dp[2] * dp[1]
        // XXXA: dp[3] * dp[0]
		withA = (withA + dp[i]*dp[n-1-i]) % 1000000007
	}
	return (dp[n] + withA) % 1000000007
}