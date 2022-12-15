func rob(nums []int) int {
    n := len(nums)
    if n < 1 {
        return 0
    }
    if n == 1 {
        return nums[0]
    }
    if n == 2 {
        if nums[0] > nums[1] {
            return nums[0]
        } else {
            return nums[1]
        }
    }
    dp := make([]int, n)
    dp[0] = nums[0]
    if nums[0] > nums[1] {
        dp[1] = nums[0]
    } else {
        dp[1] = nums[1]
    }
    for i := 2; i < n; i++ {
        a := nums[i] + dp[i - 2]
        b := dp[i - 1]
        if a > b {
            dp[i] = a
        } else {
            dp[i] = b
        }
    }
    return dp[n - 1]
}