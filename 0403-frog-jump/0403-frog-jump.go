func canCross(stones []int) bool {
	if stones[1] != 1 {
		// stone[1] must be 1, because we only can jump 1 at stone[0]
		return false
	}
	// k: value, v: value index @stones
	hash := make(map[int]int, len(stones))
	for i, v := range stones {
		hash[v] = i
	}
	// dp[i]: prev jump distances
	dp := make([]map[int]struct{}, len(stones))
	for i := range dp {
		dp[i] = make(map[int]struct{})
	}
	dp[1][1] = struct{}{}
	for i := 1; i < len(stones); i++ {
		for move := range dp[i] {
			if move-1 > 0 {
				if nextIndex := hash[stones[i]+move-1]; nextIndex > 0 {
					dp[nextIndex][move-1] = struct{}{}
				}
			}
			if nextIndex := hash[stones[i]+move]; nextIndex > 0 {
				dp[nextIndex][move] = struct{}{}
			}
			if nextIndex := hash[stones[i]+move+1]; nextIndex > 0 {
				dp[nextIndex][move+1] = struct{}{}
			}
		}
	}
	return len(dp[len(stones)-1]) > 0
}