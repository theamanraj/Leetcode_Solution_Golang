func getRow(rowIndex int) []int {    
    row := []int{1}
    ind := 0
    
    for ind < rowIndex {
        newRow := make([]int, 0, ind+2)
        newRow = append(newRow, 1)
        
        for i := 1; i < len(row); i++ {
            newRow = append(newRow, row[i]+row[i-1])
        }
        newRow = append(newRow, 1)
        row = newRow
        ind++
    }
    return row
}