func minInsertions(s string) int {
	if len(s) == 0 {
		return 0
	}
	l := len(s)
	cache := make([][]int, l)

	for i := range cache {
		cache[i] = make([]int, l)
		for j := range cache[i] {
			cache[i][j] = math.MaxInt32
		}
	}

	return dfs(s, 0, len(s)-1, &cache)
}

func dfs(s string, left, right int, cache *[][]int) int {
	if left >= right {
		return 0
	}

	if (*cache)[left][right] != math.MaxInt32 {
		return (*cache)[left][right]
	}

	if s[left] == s[right] {
		return dfs(s, left+1, right-1, cache)
	} else {
		addLeft := dfs(s, left, right-1, cache) + 1
		addRight := dfs(s, left+1, right, cache) + 1

		if addLeft > addRight {
			(*cache)[left][right] = addRight
		} else {
			(*cache)[left][right] = addLeft
		}
		return (*cache)[left][right]
	}
}