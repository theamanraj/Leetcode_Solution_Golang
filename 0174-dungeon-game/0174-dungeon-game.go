func calculateMinimumHP(dungeon [][]int) int {
    m := len(dungeon)
    n := len(dungeon[0])
    minHP := make([][]int, m + 1)
    for i := m; i > -1; i-- {
        minHP[i] = make([]int, n + 1)
        for j := n; j > -1; j-- {
            if i == m || j == n {
                minHP[i][j] = 1
            } else {
                var min int
                if i == m - 1 || (j != n - 1 && minHP[i][j + 1] < minHP[i + 1][j]) {
                    min = minHP[i][j + 1]
                } else {
                    min = minHP[i + 1][j]
                }
                step := dungeon[i][j]
                r := min - step
                if r <= 1 {
                    minHP[i][j] = 1
                } else {
                    minHP[i][j] = r
                }
            }
        }
    }
    // fmt.Println(minHP)
    return minHP[0][0]
}