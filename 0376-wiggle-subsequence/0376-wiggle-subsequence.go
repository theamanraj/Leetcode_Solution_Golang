func wiggleMaxLength(nums []int) int {
    n, res := len(nums), 2
    if n < 2 { return n }
    if nums[0] == nums[1] { res = 1 }
    preDiff := nums[0] - nums[1]
    for i := 2; i < n; i++ {
        diff := nums[i - 1] - nums[i]
        if (diff > 0 && preDiff <= 0) || (diff < 0 && preDiff >= 0) {
            res++
            preDiff = diff
        }
    }
    return res
}