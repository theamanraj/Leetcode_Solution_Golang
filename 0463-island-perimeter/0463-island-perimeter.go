func islandPerimeter(grid [][]int) int {
    
    var perimeter int
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[0]); j++ {
            if grid[i][j] == 0 {
                continue
            }
            perimeter += 4
            for _, diff := range [][]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}} {
                x, y := j + diff[0], i + diff[1]
                if y < 0 || y >= len(grid) || x < 0 || x >= len(grid[0]) {
                    continue
                }
                if grid[y][x] == 1 {
                    perimeter--
                }
            }
        }
    }
    return perimeter
}
