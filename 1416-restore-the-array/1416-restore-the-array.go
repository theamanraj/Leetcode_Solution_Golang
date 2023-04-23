func numberOfArrays(s string, k int) int {
	if s[0] == '0' {
		return 0
	}
	if k <= 9 {
		for _, c := range s {
			if int(c-'0') > k || int(c-'0') == 0 {
				return 0
			}
		}
		return 1
	}
	dp := make([]int, len(s))
	dp[0] = 1
	kWide := len(strconv.Itoa(k))
	for i := 1; i < len(s); i++ {
		if s[i] != '0' {
			// k ≥ 1
			dp[i] = dp[i-1]
		}
		if i+1 <= kWide {
			v, _ := strconv.Atoi(s[:i+1])
			if v <= k {
				dp[i]++
			}
		}
		for j := i - 2; j >= 0 && i-j <= kWide; j-- {
			if s[j+1] == '0' {
				// skip leading zero
				continue
			}
			if i-j < kWide {
				dp[i] += dp[j]
			} else {
				v, _ := strconv.Atoi(s[j+1 : i+1])
				if v <= k {
					dp[i] += dp[j]
				}
			}
		}
		dp[i] %= 1000000007
	}
	return dp[len(dp)-1]
}