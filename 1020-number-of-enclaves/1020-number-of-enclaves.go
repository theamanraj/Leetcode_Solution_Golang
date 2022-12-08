func numEnclaves(grid [][]int) int {
	var mp = make(map[int]map[int]bool)
	var finalcout int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			var flag bool
			var arr []int
			if grid[i][j] == 1 && !mp[i][j] {
				var count int
				mp = DFS(grid, mp, &count, &flag, i, j, arr)
				finalcout += count
			}
		}
	}
	return finalcout
}

func DFS(grid [][]int, mp map[int]map[int]bool, count *int, flag *bool, i, j int, arr []int) map[int]map[int]bool {
	if mp[i] == nil {
		mp[i] = make(map[int]bool)
	}
	if (grid[i][j] == 1) && (i == 0 || i == len(grid)-1 || j == 0 || j == len(grid[i])-1) {
		*flag = true
	}
	mp[i][j] = true
	arr = append(arr, 1)
	*count = *count + 1

	if i > 0 && grid[i-1][j] == 1 && !mp[i-1][j] {
		mp = DFS(grid, mp, count, flag, i-1, j, arr)
	}
	if j < len(grid[0])-1 && grid[i][j+1] == 1 && !mp[i][j+1] {
		mp = DFS(grid, mp, count, flag, i, j+1, arr)
	}
	if i < len(grid)-1 && grid[i+1][j] == 1 && !mp[i+1][j] {
		mp = DFS(grid, mp, count, flag, i+1, j, arr)
	}
	if j > 0 && grid[i][j-1] == 1 && !mp[i][j-1] {
		mp = DFS(grid, mp, count, flag, i, j-1, arr)
	}
	if *flag {
		*count = 0
	}
	return mp
}
