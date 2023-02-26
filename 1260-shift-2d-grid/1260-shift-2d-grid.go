func shiftGrid(grid [][]int, k int) [][]int {
    m, n := len(grid), len(grid[0])
    
    i := func(x int) int { return x / n }
    j := func(x int) int { return x % n }
    
    for k = k % (m*n); k > 0; k-- {
        tmp := grid[m-1][n-1]
        for ii := m*n-1; ii > 0; ii-- {
            grid[i(ii)][j(ii)] = grid[i(ii-1)][j(ii-1)]
        }
        grid[0][0] = tmp
    }
    
    return grid
}