func grayCode(n int) []int {
	l := 1
	for i := n; i > 0; i-- {
		l *= 2
	}
	rt := make([]int, l)
	rt[0] = 0

	for i, b := 1, 1; i <= n; i, b = i+1, b*2 {
		reflect(rt, i, b)
	}

	return rt
}

func reflect(n []int, j, b int) {
	for i := 0; i < b; i++ {
		n[b*2-1-i] = n[i] | 1<<uint(j-1)
	}
}