func oddCells(n int, m int, indices [][]int) int {
    countsRow := make([]int, n)
    countsCol := make([]int, m)
    for _, ind := range indices {
        countsRow[ind[0]]++
        countsCol[ind[1]]++
    }
    res := 0
    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            total := countsRow[i]+countsCol[j]
            if total % 2 == 1 {
                res++
            }
        }
    }
    return res
}