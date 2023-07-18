func minSubsequence(nums []int) []int {
    sum := 0
    res := []int{}
    tmp := 0
    for _, v := range nums {
        sum += v
    }
    sort.Slice(nums, func(a, b int) bool { return nums[a] > nums[b] })
    for _, v := range nums {
        res = append(res, v)
        tmp += v
        if tmp > sum - tmp { break }
    }
    return res
}