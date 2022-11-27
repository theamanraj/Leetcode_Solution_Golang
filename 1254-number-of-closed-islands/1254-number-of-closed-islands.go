func closedIsland(grid [][]int) int {
    
    for r:=0; r<len(grid); r++{
        for c:=0; c<len(grid[0]); c++{
            
            if r==0 || r==len(grid)-1{
                DFS(r, c, grid)
            }
            if c==0 || c==len(grid[0])-1{
                DFS(r, c, grid)
            }
        }
    }
    
    var ans int
    for r:=0; r<len(grid); r++{
        for c:=0; c<len(grid[0]); c++{
            
            if grid[r][c]==0{
                DFS(r, c, grid)
                ans++
            }
        }
    }
    
    return ans
    
}



func DFS(row int, col int, grid [][]int){
    
    if row<0 || col<0 || row>=len(grid) || col>=len(grid[0]) || grid[row][col]!=0{
        return
    }
    
    grid[row][col]=-1
    DFS(row+1, col ,grid)
    DFS(row-1, col ,grid)
    DFS(row, col+1 ,grid)
    DFS(row, col-1 ,grid)
    
}