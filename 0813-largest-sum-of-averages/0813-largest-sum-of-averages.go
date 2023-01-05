func largestSumOfAverages(A []int, K int) float64 {
	dp := make([][]float64, len(A)+1)
	for i := range dp {
		dp[i] = make([]float64, K+1)
	}
	return helper(A, dp, len(A), K)
}

func helper(nums []int, dp [][]float64, total, part int) float64 {
	if dp[total][part] != 0 {
		return dp[total][part]
	}
	if total == part {
		dp[total][part] = float64(sum(nums[:total]))
		return dp[total][part]
	}
	if part == 1 {
		dp[total][part] = calcAverage(nums[:total])
		return dp[total][part]
	}
	maxAverages := 0.0
	for i := part-1; i <= total-1; i++ {
		if curAverages := helper(nums, dp, i, part-1) + calcAverage(nums[i:total]); curAverages > maxAverages {
			maxAverages = curAverages
		}
	}
	dp[total][part] = maxAverages
	return dp[total][part]
}

func calcAverage(nums []int) float64 {
	summary := 0
	for _, v := range nums {
		summary += v
	}
	return float64(summary)/float64(len(nums))
}

func sum(values []int) int {
	summary := 0
	for _, v := range values {
		summary += v
	}
	return summary
}