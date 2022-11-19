func searchMatrix(matrix [][]int, target int) bool {
    if len(matrix) == 0{
        return false
    }
    row, col := len(matrix)-1, len(matrix[0])-1
    i , j := row , 0
    for i>=0 && j <= col{
        if matrix[i][j] == target{
            return true
        } else if matrix[i][j] > target{
            i--
        } else{
            j++
        }
    }
    return false
}