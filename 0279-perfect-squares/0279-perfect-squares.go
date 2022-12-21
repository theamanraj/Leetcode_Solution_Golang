func numSquares(n int) int {
	if n <= 0 {
		return 0
	}
	// f means i need the square sum at least
	f := make([]int, n+1)
	f[0] = 0
	for i := 1; i <= n; i++ {
		// i could be the first i-1 square sum + 1
		f[i] = f[i-1] + 1
		for j := 1; j*j <= i; j++ {
			// For each i, it must be the sum of some number (i - j*j) and
			// a perfect square number (j*j).
			f[i] = min(f[i], f[i-j*j]+1)
		}
	}
	return f[n]
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}