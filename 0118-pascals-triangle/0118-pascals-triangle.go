func generate(numRows int) [][]int {
	ret := make([][]int, numRows)

	for i := 0; i < numRows; i++ {
		arr := make([]int, i+1)
		for j := 0; j < i+1; j++ {
			if j == 0 || i == j {
				arr[j] = 1
			} else {
				arr[j] = ret[i-1][j-1] + ret[i-1][j]
			}
		}
		ret[i] = arr
	}

	return ret
}