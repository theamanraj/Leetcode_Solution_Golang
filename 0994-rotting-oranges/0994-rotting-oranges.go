func orangesRotting(grid [][]int) int {
	dirs := [][]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}

	var time, count, totalOranges int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 || grid[i][j] == 2 {
				totalOranges++
			}
		}
	}

	for true {
		var orangeRottened bool
		var visited [][]int
		time++
		for i := 0; i < len(grid); i++ {
			for j := 0; j < len(grid[0]); j++ {
				if grid[i][j] == 2 && !nodeVisisted([]int{i, j}, visited) {
					grid[i][j] = -1
					count++
					for _, dir := range dirs {
						curX := dir[0] + i
						curY := dir[1] + j
						if curX >= 0 && curX <= len(grid)-1 && curY >= 0 && curY <= len(grid[0])-1 {
							if grid[curX][curY] == 1 {
								grid[curX][curY] = 2
								orangeRottened = true
								visited = append(visited, []int{curX, curY})
							}
						}
					}
				}
			}
		}

		if !orangeRottened {
			break
		}
	}
	if count == totalOranges {
		return time - 1
	}

	return -1
}

func nodeVisisted(node []int, visited [][]int) bool {
	for _, v := range visited {
		if v[0] == node[0] && v[1] == node[1] {
			return true
		}
	}
	return false
}