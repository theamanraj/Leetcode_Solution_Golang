func updateMatrix(matrix [][]int) [][]int {
    rows := len(matrix)
    if rows == 0 { return [][]int{} }
    cols := len(matrix[0])
    if cols == 0 { return [][]int{} }
    
    res := make([][]int, rows)
    for r := 0; r < rows; r++ {
        res[r] = make([]int, cols)
    }
    
    for r := 0; r < rows; r++ {
        for c := 0; c < cols; c++ {
            if matrix[r][c] == 0 {
                res[r][c] = 0
            } else {
                left, up := -1, -1
                if r > 0 { up = res[r-1][c] }
                if c > 0 { left = res[r][c-1] }
                if up == -1 && left == -1 {
                    res[r][c] = -1
                } else if up == -1 {
                    res[r][c] = left + 1
                } else if left == -1 {
                    res[r][c] = up + 1
                } else {
                    if left < up {
                        res[r][c] = left + 1
                    } else {
                        res[r][c] = up + 1
                    }
                }
            }
        }
    }
    
    for r := rows - 1; r >= 0; r-- {
        for c := cols - 1; c >= 0; c-- {
            cur := res[r][c]
            if r < rows - 1 {
                down := res[r+1][c]
                if cur == -1 || down + 1 < cur {
                    cur = down + 1
                }
            }
            if c < cols - 1 {
                right := res[r][c+1]
                if cur == -1 || right + 1 < cur {
                    cur = right + 1
                }
            }
            res[r][c] = cur
        }
    }
    
    return res
}