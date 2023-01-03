func strangePrinter(s string) int {
    if len(s) == 0 {
        return 0
    }
    n := len(s)
    var dp = make([][]int,n)
    for k:= range dp {
        dp[k] = make([]int,n)
    }
    for i:=0;i<n;i++ {
        dp[i][i] = 1
        if i< n-1 {
            if s[i] == s[i+1] {
                dp[i][i+1] = 1
            } else {
                dp[i][i+1] = 2
            }
        }
    }
    for leng:=2;leng<n;leng++ {
        for start:=0;start+leng < n; start ++   {
            dp[start][start+leng] = leng+1
            for k:=0;k<leng;k++ {
                tmp := dp[start][start+k] + dp[start+k+1][start+leng]
                var m int
                if s[start+k] == s[start+leng] {
                    m = tmp-1 
                } else {
                    m = tmp
                }
                dp[start][start+leng] = min(dp[start][start+leng], m)
            }
        }
    }
    return dp[0][n-1]
}

func min(a,b int) int {
    if a > b {
        return b
    }
    return a
}