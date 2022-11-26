var (
    directions = [][]int{
        []int{1, 0},
        []int{-1, 0},
        []int{0, -1},
        []int{0, 1},
    }
)

func isValid(y int, x int, grid [][]byte, visited [][]bool) bool {
    if y >= len(grid) || y < 0 {
        return false
    } else if x >= len(grid[0]) || x < 0 {
        return false
    } else if grid[y][x] == '0' {
        return false
    } else if visited[y][x] == true {
        return false
    }
    
    return true
}


func numIslands(grid [][]byte) int {
    visited := make([][]bool, len(grid))
    for i := range visited {
        visited[i] = make([]bool, len(grid[i]))
    }
    numIslands := 0
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[i]); j++ {
            if grid[i][j] == '1' && !visited[i][j] {
               queue := [][]int{}
               queue = append(queue, []int{i, j})
               for (len(queue) > 0) {
                    curr := queue[0]
                    queue = queue[1:]
                    for _, direction := range directions {
                        x := curr[0] + direction[0]
                        y := curr[1] + direction[1]
                        if (isValid(x, y, grid, visited)) {
                            visited[x][y] = true
                            queue = append(queue, []int{x, y})
                        }
                    }
                }
                numIslands++
            }
        }
    }
    return numIslands
}