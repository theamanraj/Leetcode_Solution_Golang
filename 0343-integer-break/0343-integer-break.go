func max(a, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}

func integerBreak(n int) int {
    dp := make([]int, n + 1)
    dp[1] = 1
    
    
    for i := 2; i < n + 1; i++ {
        
        for j := 1; j <= i/2; j++ {
            dp[i] = max(dp[i], max(j, dp[j]) * max(i-j, dp[i-j])) 
        }
        
    }
        
    return dp[n]
    
}
