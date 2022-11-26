func canJump(nums []int) bool {
    if len(nums) == 1 { return true }
    if nums[0] == 0 { return false }
    dp := make([]bool, len(nums))
    dp[0] = true
    for i:= 0; i < len(nums); i++ {
        for j:= i+1; j < len(nums); j++ {
            //if nums[i] == 0 { continue }
            dp[j] = dp[j] || dp[i] && (nums[i] >= j-i)
            //if dp[len(nums)-1] == true { return true }
        }
    }
    return dp[len(nums)-1]
}

