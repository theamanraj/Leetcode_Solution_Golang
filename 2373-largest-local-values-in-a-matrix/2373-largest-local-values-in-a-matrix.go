func largestLocal(grid [][]int) [][]int {
    ans := make([][]int, len(grid)-2)
    for i := 0; i < len(grid)-2; i++ {
        ans[i] = make([]int, len(grid)-2)
        for j := 0; j < len(grid)-2; j++ {
            tmp := 0
            for k := j; k < j + 3; k++ {
                for l := i; l < i + 3; l++ {
                    tmp = max(tmp, grid[l][k])
                } 
            }
            ans[i][j] = tmp
        }
    }
    return ans
}

func max(i, j int) int {
    if i > j {
        return i
    }
    return j
}
