func setZeroes(matrix [][]int) {
	n := len(matrix)
	if n == 0 {
		return
	}
	m := len(matrix[0])
	clearRow := make([]int, n)
	clearColumn := make([]int, m)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] == 0 {
				clearRow[i] = 1
				clearColumn[j] = 1
			}
		}
	}
	for col, val := range clearColumn {
		if val == 1 {
			for i := 0; i < n; i++ {
				matrix[i][col] = 0
			}
		}
	}
	for row, val := range clearRow {
		if val == 1 {
			for i := 0; i < m; i++ {
				matrix[row][i] = 0
			}
		}
	}
}