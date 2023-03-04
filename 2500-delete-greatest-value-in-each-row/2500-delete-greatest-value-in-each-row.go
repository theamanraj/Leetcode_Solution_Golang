func deleteGreatestValue(grid [][]int) int {
    for _, row := range grid {
        sort.Ints(row)
    }
    ans := 0
    for i := 0; i < len(grid[0]); i++ {
        tmp := grid[0][i]
        for j := 1; j < len(grid); j++ {
            tmp = max(tmp, grid[j][i])
        }
        ans += tmp
    }
    return ans
}

func max(i, j int) int {
    if i > j {
        return i
    }
    return j
}