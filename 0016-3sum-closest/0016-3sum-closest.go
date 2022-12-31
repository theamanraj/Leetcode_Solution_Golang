func abs(a int) int {
    if a < 0 {
        return -1 * a
    }
    return a
}
func threeSumClosest(nums []int, target int) int {
    sort.Ints(nums)
    best := 1 << 30
    for i, v := range nums {
        start := i + 1
        end := len(nums) - 1
        for start < end {
            cur := nums[start] + nums[end] + v
            if abs(target - best) > abs(target - cur) {
                best = cur
            }
            if cur > target{
                end--
            } else if cur < target {
                start++
            } else {
                return target
            }
        }
    }
    return best
}