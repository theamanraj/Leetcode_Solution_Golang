func countNegatives(grid [][]int) int {
     
    m := len(grid)
    n := len(grid[0])
    
    for i := m - 1; i >= 0; i-- { 
        sort.Ints(grid[i])
    }
    
    res := 0
    for i := m - 1; i >= 0; i-- {
        for j := 0; j <= n - 1; j++ {
            if grid[i][n-1] < 0 {
                res = res + n
                break
            }
            if grid[i][j] < 0 {
                res++
            } else {
                break
            }
        }
    }
    return res
}