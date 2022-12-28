func canPartition(nums []int) bool {
    sum, targetSum := 0, 0
    for _, v := range nums {
        sum += v
    }
    
    if sum % 2 == 0 {
        targetSum = sum / 2
    } else {
        return false
    }
    
    //dp[i][j] - with first i elements, can get the targetSum j or not
    //if j >= nums[i - 1], dp[i][j] = dp[i - 1][j] || dp[i - 1][j - nums[i - 1]]
    //else dp[i][j] = dp[i - 1][j]
    
    dp := make([][]bool, len(nums) + 1)
    for i := 0; i < len(dp); i++ {
        dp[i] = make([]bool, targetSum + 1)
    }
    
    dp[0][0] = true
    
    for i := 1; i <= len(nums); i++ {
        for j := 0; j <= targetSum; j++ {
            if j >= nums[i - 1] {
                dp[i][j] = dp[i - 1][j] || dp[i - 1][j - nums[i - 1]]
            } else {
                dp[i][j] = dp[i - 1][j]
            }
            if dp[i][targetSum] == true {
                return true //return true when found targetSum
            }
        }
        
    }
    
    return false
}