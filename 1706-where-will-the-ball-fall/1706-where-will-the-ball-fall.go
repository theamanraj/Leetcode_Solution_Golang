func findBall(grid [][]int) []int {
    m := len(grid)
    n := len(grid[0])
    var locR, locC int
    var ans []int = make([]int, n)
    for i:=0;i<n;i++ {
        locR = 0
        locC = i
        for {
            if locR == m {
                ans[i] = locC
                break
            }
            if grid[locR][locC] == 1 {
                if locC == n-1 || grid[locR][locC + 1] == -1 {
                    ans[i] = -1
                    break
                }
                locR++
                locC++
            } else {
                if locC == 0 || grid[locR][locC - 1] == 1 {
                    ans[i] = -1
                    break
                }
                locR++
                locC--
            }
        }
    }
    return ans
}