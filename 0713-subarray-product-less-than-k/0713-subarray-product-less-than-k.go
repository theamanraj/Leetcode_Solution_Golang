func numSubarrayProductLessThanK(nums []int, k int) int {
    start, p := 0, 1
    var out int
    for i, v := range nums {
        p *= v
        for p >= k && start <= i {
            p /= nums[start]
            start++
        }
        out += i - start + 1
    }
    return out
}