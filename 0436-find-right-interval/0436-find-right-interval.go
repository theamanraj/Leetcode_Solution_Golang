func findRightInterval(intervals [][]int) []int {
    sorted, n := [][3]int{}, len(intervals)
    ans := make([]int, n)
    for i := 0; i < n; i++ { sorted = append(sorted, [3]int{ intervals[i][0], intervals[i][1], i }) }
    sort.Slice(sorted, func(a, b int) bool { return sorted[a][0] < sorted[b][0] })
    for i := 0; i < n; i++ { ans[i] = helper(&sorted, intervals[i][1]) }
    return ans
}

func helper(sorted *[][3]int, end int) int {
    n := len(*sorted)
    if (*sorted)[n - 1][0] < end { return -1 }
    l, r := 0, n - 1
    for l <= r {
        mid := l + (r - l) / 2
        if (*sorted)[mid][0] >= end {
            r = mid - 1
        } else {
            l = mid + 1
        }
    }
    return (*sorted)[l][2]
}