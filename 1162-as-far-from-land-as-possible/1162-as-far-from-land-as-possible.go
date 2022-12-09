func maxDistance(grid [][]int) int {
    N := len(grid)
    dx := []int{1, 0, -1, 0}
    dy := []int{0, 1, 0, -1}
    dist := make([][]int, len(grid))
    for i := range dist {
        dist[i] = make([]int, len(grid))
    }
    curr := [][]int{}
    for i := 0; i < len(grid); i ++{
        for j := 0; j < len(grid); j ++{
            if grid[i][j] == 1 {
                curr = append(curr, []int{i, j})
            }
        }
    }
    level := 0
    exist := false
    for len(curr) != 0 {
        next := [][]int{}
        level ++
        for _, c := range curr {
            for i := 0; i < 4; i ++{
                x := c[0] + dx[i]
                y := c[1] + dy[i]
                if x >= 0 && y >= 0 && x < N && y < N && grid[x][y] == 0 && dist[x][y] == 0 {
                    next = append(next, []int{x, y})
                    dist[x][y] = level
                    exist = true
                }
            }
        }
        curr = next
    }
    if !exist {
        return -1
    } else {
        return level - 1
    }
}