func max(a, b int) int {
    if (a >= b) {
        return a
    }
    return b
}

func findMaxForm(strs []string, m int, n int) int {
    var dp [][]int
    
    for i := 0; i < m + 1; i++ {
        dp = append(dp, make([]int, n + 1))
    }
    
    for _, str := range strs {
        zeros := 0
        ones := 0
        
        for _, c := range str {
            if string(c) == "0" {
                zeros++
            } else {
                ones++
            }
        }
        
        for i := m; i >= zeros; i-- {
            for j := n; j >= ones; j-- {
                dp[i][j] = max(dp[i][j], dp[i - zeros][j - ones] + 1)
            }
        }
    }
    
    return dp[m][n]
}