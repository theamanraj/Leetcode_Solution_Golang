func maximalSquare(matrix [][]byte) int {
    rows, cols := len(matrix), len(matrix[0])
    dpGrid := make([][]int, rows+1)
    for i:=0;i<rows+1;i++{
        dpGrid[i]=make([]int, cols+1)
    } 
  
    
    var maxLen int // instead of max Area, max Len is better
    for i:=1;i<rows+1;i++{
        for j:=1;j<cols+1;j++{
            if matrix[i-1][j-1]=='1'{
                dpGrid[i][j]=1+min(dpGrid[i-1][j],dpGrid[i-1][j-1],dpGrid[i][j-1])
                if dpGrid[i][j]>maxLen{
                    maxLen=dpGrid[i][j]
                }
            }             
        }
    }
    return maxLen*maxLen 
}

func min(a, b, c int)int{
    if a<=b && a<=c { return a}
    if b<=a && b<=c {return b}
    return c
}