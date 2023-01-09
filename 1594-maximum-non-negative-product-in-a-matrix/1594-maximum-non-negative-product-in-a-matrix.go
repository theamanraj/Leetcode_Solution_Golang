func max(values ...int64) int64 {
	var maxValue int64 = math.MinInt64
	for _, v := range values {
		if v > maxValue {
			maxValue = v
		}
	}
	return maxValue
}

func min(values ...int64) int64 {
	var minValue int64 = math.MaxInt64
	for _, v := range values {
		if v < minValue {
			minValue = v
		}
	}
	return minValue
}

func maxProductPath(grid [][]int) int {
	// all products is ZERO
	if grid[0][0] == 0 {
		return 0
	}

	dp := make([][][2]int64, len(grid))
	for i := range dp {
		dp[i] = make([][2]int64, len(grid[0]))
		for j := 0; j < len(dp[i]); j++ {
			// dp[i][j][0]: negative product, +1 means invalid
			// dp[i][j][1]: positive product, -1 means invalid
			dp[i][j] = [2]int64{1, -1}
		}
	}

	// initial dp[0][0]
	if grid[0][0] > 0 {
		dp[0][0][1] = int64(grid[0][0])
	} else {
		dp[0][0][0] = int64(grid[0][0])
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			helper(grid, dp, i, j)
		}
	}
	if dp[len(grid)-1][len(grid[0])-1][1] < 0 {
		return -1
	}
	return int(dp[len(grid)-1][len(grid[0])-1][1] % 1000000007)
}

func helper(grid [][]int, dp [][][2]int64, i, j int) {
	// skip 0, 0
	if i+j == 0 {
		return
	}
	if grid[i][j] == 0 {
		dp[i][j] = [2]int64{0, 0}
		return
	}
	neg := []int64{}
	pos := []int64{}
	if i > 0 {
		neg = append(neg, dp[i-1][j][0])
		pos = append(pos, dp[i-1][j][1])
	}
	if j > 0 {
		neg = append(neg, dp[i][j-1][0])
		pos = append(pos, dp[i][j-1][1])
	}
	minValue := min(neg...)
	maxValue := max(pos...)
	if grid[i][j] > 0 {
		if minValue <= 0 {
			dp[i][j][0] = int64(grid[i][j]) * minValue
		}
		if maxValue >= 0 {
			dp[i][j][1] = int64(grid[i][j]) * maxValue
		}
	} else {
		if minValue <= 0 {
			dp[i][j][1] = int64(grid[i][j]) * minValue
		}
		if maxValue >= 0 {
			dp[i][j][0] = int64(grid[i][j]) * maxValue
		}
	}
}