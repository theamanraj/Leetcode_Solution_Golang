func kInversePairs(n int, k int) int {
    
    if n == 2 {
        if k >=2 {
            return 0
        } else {
            return 1
        }
    } else if k == 0 {
        return 1
    } else if n == 1 {
        return 0
    }
    
    
    dp := make([][]int, n+1)
    for i:=0; i<=n; i++ {
        dp[i] = make([]int, k+1)
    }
    
    dp[2][0] = 1
    dp[2][1] = 1
    
    for i:=3; i<=n; i++ {
        
        dp[i][0] = 1
        
        for j:=1; j<=k; j++ {
            dp[i][j] = (dp[i][j-1]%(1e9+7) + dp[i-1][j]%(1e9+7))%(1e9+7)
            
            if j >= i {
                dp[i][j] -= dp[i-1][j-i]%(1e9+7)
                
                if dp[i][j] < 0 {
                    dp[i][j] += 1e9+7
                }
            }
        }
    }
    
    return dp[n][k]%(1e9+7)
}