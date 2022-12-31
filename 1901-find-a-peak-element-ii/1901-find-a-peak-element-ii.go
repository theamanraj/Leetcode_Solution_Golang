func findPeakGrid(mat [][]int) []int {
	return find(mat, 0, 0)
}

func find(mat [][]int, i, j int) []int {
	nextX, nextY := -1, -1
	maxValue := mat[i][j]
	if i - 1 >= 0 {
		if mat[i-1][j] > maxValue {
			maxValue = mat[i-1][j]
            nextX, nextY = i-1, j
		}
	}
	if i + 1 < len(mat) {
		if mat[i+1][j] > maxValue {
			maxValue = mat[i+1][j]
            nextX, nextY = i+1, j
		}
	}
	if j - 1 >= 0 {
		if mat[i][j-1] > maxValue {
			maxValue = mat[i][j-1]
            nextX, nextY = i, j-1
		}
	}
	if j + 1 < len(mat[0]) {
		if mat[i][j+1] > maxValue {
			maxValue = mat[i][j+1]
            nextX, nextY = i, j+1
		}
	}
	if nextX == -1 {
		return []int{i, j}
	}
	return find(mat, nextX, nextY)
}