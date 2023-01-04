func find132pattern(nums []int) bool {
    top, Max := len(nums), math.MinInt32
    for i := len(nums) - 1; i >= 0; i-- {
        if nums[i] < Max { return true }
        for top < len(nums) && nums[i] > nums[top] {
            Max = nums[top]
            top++
        }
        top--
        nums[top] = nums[i]
    }
    return false
}

func min(a, b int) int {
    if a < b { return a }
    return b
}

func max(a, b int) int {
    if a > b { return a }
    return b
}