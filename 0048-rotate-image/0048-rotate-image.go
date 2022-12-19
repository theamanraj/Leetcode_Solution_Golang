func rotate(matrix [][]int)  {
    
    n := len(matrix)-1
    
    for i:=0;i< (len(matrix)+1)/2; i++{
        for j:=0;j<(len(matrix))/2;j++{
		
            matrix[i][j], matrix[j][n-i], matrix[n-i][n-j], matrix[n-j][i] =  matrix[n-j][i], matrix[i][j], matrix[j][n-i], matrix[n-i][n-j]

        }
    }

}
