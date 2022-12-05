func matrixReshape(mat [][]int, r int, c int) [][]int {
    n, m := len(mat), len(mat[0])
    if n * m != r * c { return mat }
    ans := make([][]int, r)
    for i := range ans { ans[i] = make([]int, c) }
    for i := 0; i < n * m; i++ { ans[i / c][i % c] = mat[i / m][i % m] }
    return ans
}