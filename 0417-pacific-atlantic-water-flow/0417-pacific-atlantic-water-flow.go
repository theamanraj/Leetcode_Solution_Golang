func pacificAtlantic(heights [][]int) [][]int {
    
    rl, cl := len(heights), len(heights[0])
    
    // the cell that already flow frow four direction.
    used := make([][]int, rl)
    for i := 0; i < len(used); i++ {
        used[i] = make([]int, cl)
    }
    
    tag := 1
    
    res := [][]int{}
    for i := 0; i < rl; i++ {
        for j := 0; j < cl; j++ {
            if bfs(heights, i, j, rl, cl, used, tag) {
                res = append(res, []int{i, j})
            }
            tag++
        }
    }
    
    return res
}

var DIRS [][]int = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0,1}}
                 
func bfs(g [][]int, ri int, ci int, rl int, cl int, used [][]int, tag int) bool {
    
    queue := list.New()
    queue.PushFront([]int{ri, ci})
    used[ri][ci] = tag
    
    mask := 0
    for queue.Len() > 0 {
        xy := queue.Back().Value.([]int)
        queue.Remove(queue.Back())
        x, y := xy[0], xy[1]
        if x == 0 || y == 0 {
            mask |= 1
        }
        
        if x == rl - 1 || y == cl - 1{
            mask |= 2
        }
        
        if mask == 3 {
            return true
        }
        
        for _, d := range DIRS {
            x1, y1 := x + d[0], y + d[1]
            // neighboring cell
            if x1 >= 0 && x1 < rl && y1 >= 0 && y1 < cl && used[x1][y1] != tag && g[x1][y1] <= g[x][y] {
                used[x1][y1] = tag
                queue.PushFront([]int{x1, y1})
            }
        }
    }
    
    return false
}