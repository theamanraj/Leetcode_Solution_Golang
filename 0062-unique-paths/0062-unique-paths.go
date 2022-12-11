func uniquePaths(m int, n int) int {
    if m == 1 {
        return 1
    }
    if n == 1 {
        return 1
    }
    step := [][]int{}
    for i := 0; i < m; i++ {
        step = append(step, make([]int, n))
    }
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if i == 0  || j == 0 {
                step[i][j] = 1
                continue
            }
            step[i][j] = step[i-1][j] + step[i][j-1]
        }
    }
    return step[len(step)-1][len(step[0])-1]
}