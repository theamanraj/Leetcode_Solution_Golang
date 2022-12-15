func numTrees(n int) int {
dp := make([]int, n+1)
dp[0] = 1
dp[1] = 1
for index := 2; index <= n; index++ {
for i, j := index-1, 0; i >= 0; i, j = i-1, j+1 {
dp[index] += dp[i] * dp[j]
}
}
return dp[n]
}