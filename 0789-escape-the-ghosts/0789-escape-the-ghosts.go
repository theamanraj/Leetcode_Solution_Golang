func escapeGhosts(g [][]int, tgt []int) bool {
    dist := abs(tgt[0]) + abs(tgt[1])
    for _, v := range g {
        cur := abs(v[0] - tgt[0]) + abs(v[1] - tgt[1])
        if cur <= dist {
            return false
        }
    }
    return true
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}