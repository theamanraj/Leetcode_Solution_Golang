func maxSumSubmatrix(matrix [][]int, k int) int {
    m := len(matrix)
    n := len(matrix[0])
    dp := make([][]int, m)
    for i := 0; i < m; i++ {
        dp[i] = make([]int, n)
    }
    dp[0][0] = matrix[0][0]
    for i := 1; i < m; i++ {
        dp[i][0] = dp[i - 1][0] + matrix[i][0]
    }
    for i := 1; i < n; i++ {
        dp[0][i] = dp[0][i - 1] + matrix[0][i]
    }
    
    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            dp[i][j] = dp[i-1][j] + dp[i][j-1] - dp[i-1][j-1] + matrix[i][j]
        }
    }
    res := -2<<32
    A, B, C, D := 0, 0, 0, 0
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            for x := i; x < m; x++ {
                for y := j; y < n; y++ {
                    A = dp[x][y]
                    if j - 1 >= 0 {
                        B = dp[x][j - 1]
                    } else {
                        B = 0
                    }
                    if i - 1 >= 0 {
                        C = dp[i - 1][y]
                    } else {
                        C = 0
                    }
                    
                    if i - 1 >= 0 && j - 1 >= 0 {
                        D = dp[i - 1][j - 1]
                    } else {
                        D = 0
                    }
                    area := A - B - C + D
                    if area <= k {
                        res = max(res, area)
                    }
                }
            }
        }
    }
    return res
}

func max(a, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}