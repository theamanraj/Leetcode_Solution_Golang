func isValidSudoku(board [][]byte) bool {
     row := make(map[string]bool)
     col := make(map[string]bool)
     block := make(map[string]bool)
    
    for i :=0;i<9;i++{
        for j:=0;j<9;j++{
           current_val :=string(board[i][j])
           if current_val =="." { continue}
            ok1 := row[string(i)+"_"+current_val]
            ok2 := col[string(j)+"_"+current_val]
            ok3 := block[string(i/3) + "-" + string(j/3)+current_val] 
            if ok1 || ok2 || ok3{
                return false
            }else {
                row[string(i)+"_"+current_val] = true
                col[string(j)+"_"+current_val]= true
                block[string(i/3) + "-" + string(j/3)+current_val] = true
            }
        }
    }
    return true
}