func mergeSimilarItems(items1 [][]int, items2 [][]int) [][]int {
    m := make(map[int]int)
    for _, items := range items1 {
        if v, ok := m[items[0]]; ok {
            m[items[0]] = v + items[1]
        } else {
            m[items[0]] = items[1]
        }
    }
    for _, items := range items2 {
        if v, ok := m[items[0]]; ok {
            m[items[0]] = v + items[1]
        } else {
            m[items[0]] = items[1]
        }
    }
    ans := make([][]int, 0, len(m))
    for k, v := range m {
        t := []int{k, v}
        ans = append(ans, t)
    }
    sort.Slice(ans, func(i, j int) bool { return ans[i][0] < ans[j][0] })
    return ans
}