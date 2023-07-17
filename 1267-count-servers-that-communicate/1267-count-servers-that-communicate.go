func countServers(grid [][]int) int {
    m := len(grid)
    n := len(grid[0])
    mark := make([][]int, m)
    for i := 0; i < m; i ++{
        mark[i] = make([]int, n)
    }
    for i := 0; i < m; i ++{
        sum := 0
        for j := 0; j < n; j ++{
            sum += grid[i][j]
        }
        if sum > 1 {
            for j := 0; j < n; j ++{
                if grid[i][j] == 1 {
                    mark[i][j] = 1
                }
            }
        }
    }    
    for j := 0; j < n; j ++{
        sum := 0
        for i := 0; i < m; i ++{
            sum += grid[i][j]
        }
        if sum > 1 {
            for i := 0; i < m; i ++{
                if grid[i][j] == 1 {
                    mark[i][j] = 2
                }
            }
        } 
    }    
    total := 0
    for i := 0; i < m; i ++{
        for j := 0; j < n; j ++{
            if mark[i][j] != 0 {
                total ++
            }
        }
    }
    return total
}