func isOverlapping(e1, e2 []int) bool {
    if e1[1] < e2[0] || e2[1] < e1[0] {
        return false
    }
    return true
}

func min(a, b int) int {
    if a > b { return b }
    return a
}

func max(a, b int) int {
    if a < b { return b }
    return a
}

func getOverlappingRange(e1, e2 []int) []int{
    if e1[0] == e2[1] {
        return []int{e1[0], e1[0]}
    } else if e1[1] == e2[0] {
        return []int{e1[1], e1[1]}
    }
    return []int{max(e1[0], e2[0]), min(e1[1], e2[1])}
}

func intervalIntersection(A [][]int, B [][]int) [][]int {
    i, j := 0, 0
    res := make([][]int, 0)
    for i < len(A) && j < len(B) {
        e1, e2 := A[i], B[j]
        if isOverlapping(e1, e2) == true {
            res = append(res, getOverlappingRange(e1, e2))
        }
        if e2[1] >= e1[1] {
            i = i + 1
        } else {
            j = j + 1
        }
    }
    return res
}