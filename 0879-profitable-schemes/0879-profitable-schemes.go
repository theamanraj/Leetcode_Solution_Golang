func profitableSchemes (G int, P int, group []int, profit []int) int {
    MOD := 1000000007
    dp := make([][][]int, 2)
    for i := 0; i < 2; i++ {
        dp[i] = make([][]int, G + 1) 
        for j := 0; j < len(dp[i]); j++ {
            dp[i][j] = make([]int, P + 1)
        }
    }
    dp[0][0][0] = 1

    for i := 1; i <= len(group); i++ {
        for j := 0; j <= G; j++ {
            for k := 0; k <= P; k++ {
                dp[i%2][j][k] = dp[(i-1)%2][j][k] % MOD
                if j - group[(i-1)] >= 0 {
                    preProfit := max(k - profit[(i-1)], 0)
                    dp[i%2][j][k] += dp[(i-1)%2][j-group[(i-1)]][preProfit] % MOD
                }
            }
        }
    }
    
    res := 0
    for j := 0; j <= G; j++ {
        res += dp[len(group)%2][j][P]
    }
    return res % MOD
}

func max(a, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}