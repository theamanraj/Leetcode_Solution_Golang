func smallestRangeII(A []int, K int) int {
    if len(A) == 1 {
        return 0
    }
    sort.Ints(A) 
    res := A[len(A) - 1] - A[0] //default value, for the edge case 
    P1 := A[0] + K
    P4 := A[len(A) - 1] - K
    for i := 0; i < len(A) - 1; i++ {
        P2 := A[i] + K
        P3 := A[i + 1] - K
        res = min(res, max(P2, P4) - min(P1, P3))
    }
    return res
}

func min(a, b int) int {
    if a < b {
        return a
    } else {
        return b
    }
}

func max(a, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}