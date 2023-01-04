func findMinArrowShots(points [][]int) int {
    sort.Slice(points, func(i, j int) bool {
        return points[i][0] < points[j][0] || (points[i][0] == points[j][0] && points[i][1] < points[j][1])
    })
    ans := 0
    cur := 0
    for _, v := range points {
        if cur == 0 {
            cur = v[1]
            ans += 1
            continue
        }
        if cur < v[0] {
            cur = v[1]
            ans += 1
            continue
        }
        if cur > v[1] {
            cur = v[1]
        }
    }
    return ans
}
