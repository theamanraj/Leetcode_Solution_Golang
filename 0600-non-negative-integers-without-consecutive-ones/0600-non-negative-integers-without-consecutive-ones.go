func findIntegers(num int) int {
	dp := make([]int, 32)
	dp[0] = 1
	dp[1] = 2
	dp[2] = 3
	for i := 3; i < len(dp); i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return helper(num, dp)
}

func helper(num int, dp []int) int {
	if num <= 1 {
		return num+1
	}
	width := 0
	numCopy := num
	for numCopy != 0 {
		width++
		numCopy >>= 1
	}
	if width == 2 {
		return 3
	}
	total := dp[width-1]
	sentinel := 1<<(width-1) + 1<<(width-2) - 1
	if num >= sentinel {
		total += dp[width-2]
	} else {
		total += helper(num-1<<(width-1), dp)
	}
	return total
}