type NumMatrix struct {
    dp [][]int
}

func Constructor(matrix [][]int) NumMatrix {
    var mat NumMatrix
    
    if len(matrix) == 0 || len(matrix[0]) == 0 {
        return mat
    }

    for i := 0; i < len(matrix); i++ {
        mat.dp = append(mat.dp, make([]int, len(matrix[0]) + 1))
    }
    
    for r := 0; r < len(matrix); r++ {
        for c := 0; c < len(matrix[0]); c++ {
            mat.dp[r][c + 1] = mat.dp[r][c] + matrix[r][c]
        }
    }
    
    return mat
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
    sum := 0
    
    for row := row1; row <= row2; row++ {
        sum += this.dp[row][col2 + 1] - this.dp[row][col1]
    }
    
    return sum
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */