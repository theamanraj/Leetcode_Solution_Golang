func transpose(A [][]int) [][]int {
    row:=len(A[0])
    col:=len(A)
   
    matrix:=make([][]int,0,row)
     for j:=0;j<len(A[0]);j++{
          strArray:=make([]int,0,col)
    for i:=0;i<len(A);i++{
        strArray=append(strArray,A[i][j])
    }
         matrix=append(matrix,strArray)
    }
    return matrix
}