func findLongestChain(pairs [][]int) int {
    sort.Slice(pairs, func(a, b int) bool { return pairs[a][1] < pairs[b][1] })
    cur := math.MinInt32
    res := 0
    for _, v := range pairs {
        if cur < v[0] {
            cur = v[1]
            res++
        }
    }
    return res
}