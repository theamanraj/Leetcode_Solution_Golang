func maxCoins(nums []int) int {
    
    //create a new slice val to store the input nums
    val := make([]int, 0)
    val = append(val, 1) //nums[-1]
    val = append(val, nums...)
    val = append(val, 1) //nums[n]
    n := len(val)
    
    //dp[i][j] - the max coins can collect between position i and j (both exclusive)
    dp := make([][]int, len(nums) + 2)
    for i := 0; i < len(dp); i++ {
        dp[i] = make([]int, len(dp) + 2)
    }
    
    //iterate from length l == 3
    for l := 3; l <= n; l++ {
        //i is the left position (exclusive)
        for i := 0; i <= n - l; i++ {
            j := i + l - 1 //j is the right position (exclusive)
            for k := i + 1; k < j; k++ { 
                sum := dp[i][k] + dp[k][j] + val[i] * val[k] * val[j]
                dp[i][j] = max(dp[i][j], sum)
            }
        }
    }
    return dp[0][n - 1]
}

func max(a, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}