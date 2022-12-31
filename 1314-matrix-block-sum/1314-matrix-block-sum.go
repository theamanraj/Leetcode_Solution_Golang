func matrixBlockSum(mat [][]int, k int) [][]int {
    // prefix matrix
    
    // 1. build prefix matrix
    for i := range mat {
        for j := range mat[i] {
            left := 0
            if j > 0 {
                left = mat[i][j-1]
            }
            up := 0
            if i > 0 {
                up = mat[i-1][j]
            }
            diag := 0
            if i > 0 && j > 0 {
                diag = mat[i-1][j-1]
            }
            mat[i][j] += left + up - diag
        }
    }
    
    // 2. build res
    n, m := len(mat), len(mat[0])
    res := make([][]int, len(mat))
    for i := range res {
        res[i] = make([]int, len(mat[i]))
        for j := range res[i] {
            
            // 2.1 get the big encompassing mat
            bigX := min(i+k, n-1)  // last legal x
            bigY := min(j+k, m-1)  // last legal y
            bigMat := mat[bigX][bigY]
            
            // 2.2 get the left mat to subtract
            y := max(j-k-1, -1)
            var leftMat int
            if y > -1 {
                leftMat = mat[bigX][y]
            }
            // 2.3 get the upper mat to subtract
            x := max(i-k-1, -1)
            var upMat int
            if x > -1 {
                upMat = mat[x][bigY]
            }
            
            // 2.4 get the upper left mat to add back, since it is subtracted twice
            var upLeftMat int
            if x != -1 && y != -1 {
                upLeftMat = mat[x][y]
            }
            
            // 2.5 calculate size
            res[i][j] = bigMat - leftMat - upMat + upLeftMat
        }
    }
    
    return res
}

func min(a, b int) int {
    if a < b {
        return a
    }
    
    return b
}

func max(a, b int) int {
    if a > b {
        return a
    }
    
    return b
}