func maxAreaOfIsland(grid [][]int) int {
    maxArea := 0
    
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[0]); j++ {
            if grid[i][j] == 1 {
                area := getArea(grid, i, j)
                if area > maxArea {
                    maxArea = area
                }
            }
        }
    }
    
    return maxArea
}

func getArea(grid [][]int, row, colum int) int {
    if grid[row][colum] == 0 {
        return 0
    }
    
    area := 1
    
    grid[row][colum] = 0
    
    if row > 0 {
        area += getArea(grid, row - 1, colum)
    }
    if row < len(grid) - 1 {
        area += getArea(grid, row + 1, colum)
    }
    if colum > 0 {
        area += getArea(grid, row, colum - 1)
    }
    if colum < len(grid[0]) - 1 {
        area += getArea(grid, row, colum + 1)
    }
    
    return area
}