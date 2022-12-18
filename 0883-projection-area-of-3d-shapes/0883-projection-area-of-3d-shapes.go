func projectionArea(grid [][]int) int {
    var res int

	lenRow := len(grid)
	lenCol := len(grid[0])

	for i := 0; i < lenRow; i ++ {
		maxX, maxY := 0,0
		for j := 0; j < lenCol; j ++ {
			// z
			if grid[i][j] > 0 {
				res ++
			}

            // x
			if grid[i][j] > maxX {
				maxX = grid[i][j]
			}

            // y
			if grid[j][i] > maxY {
				maxY = grid[j][i]
			}
		}

		res = res + maxX + maxY
	}

	return res
}