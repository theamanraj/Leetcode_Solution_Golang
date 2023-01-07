func findDiagonalOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	mlen, nlen := len(matrix), len(matrix[0])

	var res []int
	i, j, goUp := 0, 0, false
	for {
		res = append(res, matrix[i][j])
		if i == mlen-1 && j == nlen-1 {
			break
		}
		if !goUp {
			if i > 0 && j < nlen-1 {
				i, j = i-1, j+1
			} else if j < nlen-1 {
				j, goUp = j+1, !goUp
			} else {
				i, goUp = i+1, !goUp
			}
			continue
		}
		if i < mlen-1 && j > 0 {
			i, j = i+1, j-1
		} else if i == mlen-1 {
			j, goUp = j+1, !goUp
		} else {
			i, goUp = i+1, !goUp
		}
	}
	return res
}