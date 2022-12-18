func minOperations(nums []int) int {
    var (
        prev = 0
        res = 0
    )

    for _, cur := range nums {
        if cur <= prev {
            prev += 1
            res += prev - cur
        } else {
            prev = cur
        }
    }

    return res
}