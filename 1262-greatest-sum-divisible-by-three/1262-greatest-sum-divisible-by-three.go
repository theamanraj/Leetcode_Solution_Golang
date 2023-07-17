func maxSumDivThree(nums []int) int {
    dp := make([]int, 3)
    for _, v := range nums {
        tmp := make([]int, 3)
        copy(tmp, dp)
        for i := 0; i < 3; i++ {
            dp[(v + tmp[i]) % 3] = max(dp[(v + tmp[i]) % 3], v + tmp[i])
        }
    }
    return dp[0]
}

func max(a, b int) int {
    if a > b { return a }
    return b
}