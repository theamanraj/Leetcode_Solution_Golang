func findMaxLength(nums []int) int {
    ans, sum, mem := 0, 0, make(map[int]int)
    mem[0] = -1
    for i, v := range nums {
        if v == 1 {
            sum++
        } else {
            sum--
        }
        if index, ok := mem[sum]; ok {
            ans = max(ans, i - index)
        } else {
            mem[sum] = i
        }
    }
    return ans
}

func max(a, b int) int {
    if a > b { return a }
    return b
}