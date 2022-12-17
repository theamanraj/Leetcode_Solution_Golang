func splitArray(nums []int, m int) int {
    var max, total int = 0, 0
    for _, v := range nums {
        if max < v { max = v }
        total += v
    }
    lo, hi := max, total
    for lo <= hi {
        mid := (lo + hi) / 2
        res := canPartition(nums, mid)
        if res <= m {
            hi = mid - 1 
        } else {
            lo = mid + 1
        }
    }
    return lo
}

func canPartition(n []int, goal int) int {
    var cur, p int
    for _, v := range n {
        cur += v
        if cur > goal {
            cur = v
            p++
        }
    }
    return p + 1
}