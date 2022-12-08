func countSubIslands(grid1 [][]int, grid2 [][]int) int {
    m, n := len(grid2), len(grid2[0])
    var subIslands int
    
    for i := 0; i< m; i++ {
        for j := 0; j < n; j++ {
            if grid2[i][j] == 1 && dfs(i, j, grid1, grid2) {
                subIslands++
            }
        }
    }
    return subIslands
}

func dfs(i, j int, grid1, grid2 [][]int) bool {
    if i < 0 || j < 0 || i == len(grid2) || j == len(grid2[0]) || grid2[i][j] == 0 {
        return true
    }
	
    isSubIsland := true
	
    if grid1[i][j] == 0 {
        isSubIsland = false
    }
	
    grid2[i][j] = 0
    
    isSubIsland = dfs(i+1, j, grid1, grid2) && isSubIsland
    isSubIsland = dfs(i-1, j, grid1, grid2) && isSubIsland
    isSubIsland = dfs(i, j+1, grid1, grid2) && isSubIsland
    isSubIsland = dfs(i, j-1, grid1, grid2) && isSubIsland
    
    return isSubIsland
}