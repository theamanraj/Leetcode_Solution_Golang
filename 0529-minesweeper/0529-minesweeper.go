func dfs(board [][]byte, r int, c int) {
    
    if r >= len(board) || c >= len(board[0]) || r < 0 || c < 0 {
        return
    }
    
    if board[r][c] != 'E' {
        return
    }
    
    dirs := [][]int{{0,1},{0,-1},{1,0},{-1,0},{1,1},{1,-1},{-1,1},{-1,-1}}

    totalMines := 0
    for _, dir := range dirs {
        nr := r + dir[0]
        nc := c + dir[1]
        if nr < len(board) && nc < len(board[0]) && nr >= 0 && nc >= 0 {
            if board[nr][nc] == 'M' {
                totalMines++
            }
        }
    }
    
    if totalMines != 0 {
        board[r][c] = byte(totalMines) + '0'
        return
    }
    
    board[r][c] = 'B'
    for _, dir := range dirs {
        nr := r + dir[0]
        nc := c + dir[1]
        if nr < len(board) && nc < len(board[0]) && nr >= 0 && nc >= 0 {
            dfs(board, nr, nc)
        }
    }
    
}

func updateBoard(board [][]byte, click []int) [][]byte {
    
    row := click[0]
    col := click[1]
    
    if board[row][col] == 'M' {
        board[row][col] = 'X'
        return board
    }
    
    dfs(board, row, col)
    
    return board
    
}