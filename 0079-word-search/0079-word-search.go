func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])

	var dfs func(r, c, i int) bool
	dfs = func(r, c, i int) bool {
		if i >= len(word) {
			return true
		}
		if r < 0 || r >= m || c < 0 || c >= n ||
			board[r][c] != word[i] || board[r][c] == '*' {
			return false
		}
		visited := board[r][c]
		board[r][c] = '*'
		r0, r1 := dfs(r-1, c, i+1), dfs(r+1, c, i+1)
		c0, c1 := dfs(r, c-1, i+1), dfs(r, c+1, i+1)
		board[r][c] = visited
		return r0 || r1 || c0 || c1
	}

	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			if dfs(r, c, 0) {
				return true
			}
		}
	}
	return false
}