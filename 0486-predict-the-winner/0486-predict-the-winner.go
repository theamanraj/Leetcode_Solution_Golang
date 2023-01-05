func PredictTheWinner(nums []int) bool {
	dp := make([][]int, len(nums))
	for i := range dp {
		dp[i] = make([]int, len(nums))
	}
	for i := range dp {
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	score := helper(0, len(nums)-1, dp, nums)
	total := 0
	for _, v := range nums {
		total += v
	}
	return score * 2 >= total
}

func helper(i, j int, dp [][]int, nums []int) int {
	if dp[i][j] != -1 {
		return dp[i][j]
	}
	if i == j {
		dp[i][j] = nums[i]
		return dp[i][j]
	}
	if j - i == 1 {
		dp[i][j] = max(nums[i], nums[j])
		return dp[i][j]
	}
	a := helper(i+2, j, dp, nums)
	b := helper(i+1, j-1, dp, nums)
	c := helper(i, j-2, dp, nums)
	dp[i][j] = max(nums[i] + min(a, b), nums[j] + min(b, c))
	return dp[i][j]
}

func max(values ...int) int {
	maxValue := math.MinInt32
	for _, v := range values {
		if v > maxValue {
			maxValue = v
		}
	}
	return maxValue
}

func min(values ...int) int {
	minValue := math.MaxInt32
	for _, v := range values {
		if v < minValue {
			minValue = v
		}
	}
	return minValue
}