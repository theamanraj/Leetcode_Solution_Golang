func minSubArrayLen(target int, nums []int) int {
    ans := math.MaxInt64
    c := 0
    le := 0
    for i, v := range nums {
        c += v
        for c >= target && le <= i {
            if i - le + 1 < ans {
                ans = i - le + 1
            }
            c -= nums[le]
            le++
        }
    }
    if ans == math.MaxInt64 {
        return 0
    }
    return ans
}