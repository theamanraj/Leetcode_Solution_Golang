func numberOfArithmeticSlices(nums []int) int {
    res := 0
    //the number of arithmetic subsquences when last element is nums[i] and diff is d
    //as the number of d values are huge, we can maintain it in a map
    dp := make([]map[int]int, len(nums)) 
    for i, x := range nums {
        dp[i] = make(map[int]int)
        for j, y := range nums[:i] {
            d := x - y
            res += dp[j][d] //dp[j][d] can be added to result at this moment, as nums[i] will be added after nums[j]
            dp[i][d] += dp[j][d] + 1
        }
    }
    return res
}