func tictactoe(moves [][]int) string {
    
    row := make([]int, 3)
    col := make([]int, 3)
    dia := 0
    dedia := 0
    
    for i, v := range moves {
        t := 1
        if i % 2 == 1 {
            t = -1
        }
        row[v[0]] += t
        col[v[1]] += t
        if v[0] == v[1] {
            dia += t
        }
        if v[0] == (3 - 1 - v[1]) {
            dedia += t
        }
        
        if row[v[0]] == 3 || col[v[1]] == 3 || dia == 3 || dedia == 3 {
            return "A"
        }
        
        if row[v[0]] == -3 || col[v[1]] == -3 || dia == -3 || dedia == -3 {
            return "B"
        }
    }
    
    if len(moves) == 9 {
        return "Draw"
    }
    return "Pending"
}