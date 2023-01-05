func maxIncreaseKeepingSkyline(grid [][]int) int {
	var sum int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			max1 := -1
			max2 := -1
			for k := 0; k < len(grid[j]); k++ {
				if max1 < grid[i][k] {
					max1 = grid[i][k]
				}
			}
			for m := 0; m < len(grid[j]); m++ {
				if max2 < grid[m][j] {
					max2 = grid[m][j]
				}
			}
			x := Maxoftwo(max1, max2)

			if x > grid[i][j] {
				sum = sum + (x - grid[i][j])
				grid[i][j] = x
			}

		}
	}
	return sum

}
func Maxoftwo(max1, max2 int) int {
	if max1 < max2 {
		return max1
	} else {
		return max2
	}
}