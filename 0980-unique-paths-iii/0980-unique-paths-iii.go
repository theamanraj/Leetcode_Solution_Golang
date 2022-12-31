func uniquePathsIII(grid [][]int) int {
    if len(grid) == 0 {
        return 0
    }
    
    for i, row := range grid {
        for j, cell := range row{
            if cell == 1 {
                visited := make([][]int, len(grid))
                for a:=0;a<len(grid);a++{
                    visited[a] = make([]int, len(grid[0]))
                }
                
                ans := 0
                getPaths(grid, visited, i, j, &ans)
                return ans
            }
        }
    }
    return 0
}
var directions = [][]int{{1,0},{-1,0},{0,1},{0,-1}}

func getPaths(grid [][]int, visited [][]int, i, j int, ans * int)  {
    if i<0 || i >= len(grid) || j <0 || j >=len(grid[0]) {
        return
    }
    if visited[i][j] == 1 || grid[i][j] == -1{
        return
    }
    if grid[i][j] == 2 {
        for i:=0;i<len(grid);i++{
            for j:=0;j<len(grid[0]);j++{
                if grid[i][j] == 0 && visited[i][j] != 1 {
                    return
                } else {
                    
                }
            }
        }
        *ans += 1
        return
    }
    
    visited[i][j] = 1
    for _, dir := range directions {
        x := i+dir[0]
        y := j+dir[1]
        getPaths(grid, visited,x, y, ans)
        
    }
    visited[i][j] = 0
}