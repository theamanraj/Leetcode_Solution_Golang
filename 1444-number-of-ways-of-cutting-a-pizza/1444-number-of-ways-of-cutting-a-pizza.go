func ways(pizza []string, k int) int {
	rows, cols := len(pizza), len(pizza[0])
	apples := make([][]int, rows+1)
	for row := 0; row < rows+1; row++ {
		apples[row] = make([]int, cols+1)
	}
	dp := make([][][]int, k)
	for remain := 0; remain < k; remain++ {
		dp[remain] = make([][]int, rows)
		for i := 0; i < rows; i++ {
			dp[remain][i] = make([]int, cols)
		}
	}
	for row := rows - 1; row >= 0; row-- {
		for col := cols - 1; col >= 0; col-- {
			if apples[row] == nil {
				apples[row] = make([]int, cols+1)
			}
			if pizza[row][col] == 'A' {
				apples[row][col]++
			}
			apples[row][col] = apples[row][col] + apples[row][col+1] + apples[row+1][col] - apples[row+1][col+1]
			if apples[row][col] > 0 {
				dp[0][row][col] = 1
			}
		}
	}

	mod := 1000000007
	for remain := 1; remain < k; remain++ {
		for row := 0; row < rows; row++ {
			for col := 0; col < cols; col++ {
				for i := row + 1; i < rows; i++ {
					if apples[row][col] > apples[i][col] {
						dp[remain][row][col] = (dp[remain][row][col] + dp[remain-1][i][col]) % mod
					}
				}
				for j := col + 1; j < cols; j++ {
					if apples[row][col] > apples[row][j] {
						dp[remain][row][col] = (dp[remain][row][col] + dp[remain-1][row][j]) % mod
					}
				}
			}
		}
	}
	return dp[k-1][0][0]
}
