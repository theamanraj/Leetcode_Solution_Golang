const (
    dead = 0 // was dead and should be dead or was dead and is not checked yet
    alive = 1 // was alive and should be alive or was alive and is not checked yet
    shouldDie = 2 // was alive and should die
    shouldLive = 3 // was dead and should be alive
)

func gameOfLife(board [][]int)  {
    for i := 0; i < len(board); i++ {
        for j := 0; j < len(board[0]); j++ {
            setNextState(board, i, j)
        }
    }

    clean(board)
}

func setNextState(board [][]int, i, j int) {
    diffs := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1}}

    aliveNeighbours := 0
    for _, d := range diffs {
        if aliveNeighbours > 3 {
            break
        }

        ic, jc := i + d[0], j + d[1]
        if ic < 0 || ic >= len(board) || 
            jc < 0 || jc >= len(board[0]) {
                continue
            }

        if board[ic][jc] == alive || board[ic][jc] == shouldDie {
            aliveNeighbours++
        }
    }

    if board[i][j] == dead  {
        if aliveNeighbours == 3 {
            board[i][j] = shouldLive
        }
        return  
    }

    if aliveNeighbours >= 2 && aliveNeighbours <= 3 {
        return
    }

    board[i][j] = shouldDie
}

func clean(board [][]int) {
    for i := 0; i < len(board); i++ {
        for j := 0; j < len(board[0]); j++ {
            if board[i][j] == shouldDie {
                board[i][j] = dead
            }
            if board[i][j] == shouldLive {
                board[i][j] = alive
            }
        }
    }
}