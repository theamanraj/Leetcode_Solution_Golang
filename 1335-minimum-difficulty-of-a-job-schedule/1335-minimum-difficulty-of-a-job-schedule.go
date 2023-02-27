func minDifficulty(jd []int, d int) int {
	l := len(jd)
	if l < d {
		return -1
	}
	me := make([][]int, l)
	for i := range me {
		me[i] = make([]int, d+1)
		for j := range me[i] {
			me[i][j] = -1
		}
	}
	var dp func(int, int) int
	dp = func(i, j int) int {
		if j == 0 {
			if i >= l {
				return 0
			}
			return -1
		}
		if i >= l {
			return -1
		}
		if me[i][j] != -1 {
			return me[i][j]
		}
		m := 0
		r := 9999999
		for ni := i; ni < l; ni++ {
			if jd[ni] > m {
				m = jd[ni]
			}
			v := dp(ni+1, j-1)
			if v == -1 {
				continue
			}
			v += m
			if v < r {
				r = v
			}
		}
		me[i][j] = r
		return r
	}
	return dp(0, d)
}
