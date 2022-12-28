func countBattleships(board [][]byte) int {
    result := 0
    for i, yBoard := range board {
        for j, cell := range yBoard {
            if cell == 'X' {
                if i - 1 >= 0 && board[i - 1][j] == 'X' {
                    continue
                } 
                if j - 1 >= 0 && board[i][j - 1] == 'X' {
                    continue
                }
                result++
            }
        }
    }
    
    return result
}