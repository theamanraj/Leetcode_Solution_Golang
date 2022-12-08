var DIR [][]int = [][]int{
    {0, 1},
    {0, -1},
    {-1, 0},
    {1, 0},
    {1, 1},
    {-1, 1},
    {1, -1},
    {-1, -1},
}

func shortestPathBinaryMatrix(grid [][]int) int {        
    if grid[0][0] == 1{
        return -1
    }
    n := len(grid)   
    ans := 1 
    q := make([]pair, 0)
    q = append(q, pair{0,0})
    grid[0][0] = 1
    for len(q) > 0{        
        temp := make([]pair, 0)
        for _,v := range q{            
            if v.i == n - 1 && v.j == n - 1{
                return ans
            }
            for _, d := range DIR{
                x := v.i + d[0]
                y := v.j + d[1]
                if x >= 0 && y >= 0 && x < n && y < n && grid[x][y] == 0{
                    grid[x][y] = 1
                    temp = append(temp, pair{x,y})                    
                }
            }
        }        
        q = temp
        ans++
    }
    return -1
}

type pair struct{
    i int
    j int
}