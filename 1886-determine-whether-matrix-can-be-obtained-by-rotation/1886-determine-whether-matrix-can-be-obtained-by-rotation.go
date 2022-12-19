func findRotation(mat [][]int, target [][]int) bool {
    n := len(mat) 
    grid := make([][]int, n) 
    for i := 0; i < n; i++ {
        grid[i] = make([]int, n) 
        for j := 0; j < n; j++ {
            grid[i][j] = mat[i][j] 
        }
    }
    if equals(mat, target) { return true }
    count := 4 
    for count != 0 {
        for i := 0; i < n; i++ {
            for j := 0; j < n; j++ {
                grid[j][i] = mat[n - i - 1][j] 
            }
        }
        if equals(grid, target) {
            return true 
        }
        for i := 0; i < n; i++ {
            for j := 0; j < n; j++ {
                mat[i][j] = grid[i][j] 
            }
        }
        count-- 
    }
    return false 
}

func equals(matrix, grid [][]int) bool {
    n := len(matrix) 
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] != matrix[i][j] {
                return false 
            }
        }
    }
    return true 
}
