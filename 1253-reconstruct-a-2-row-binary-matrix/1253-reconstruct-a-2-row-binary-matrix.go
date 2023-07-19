func reconstructMatrix(upper int, lower int, colsum []int) [][]int {

	var (
		uc, lc = upper, lower
        res    = make([][]int, 2)
	)
    res[0], res[1] = make([]int, len(colsum)), make([]int, len(colsum))
	for i := range colsum {
		switch colsum[i] {
		case 2:
			res[0][i] = 1
			res[1][i] = 1
			uc--
			lc--
		case 1:
			switch {
			case uc >= lc:
				res[0][i] = 1
				res[1][i] = 0
				uc--
			case lc > uc:
				res[0][i] = 0
				res[1][i] = 1
				lc--
			}
		}
	}
    if uc != 0 || lc != 0 {
        return [][]int{}
    }
	return res
}