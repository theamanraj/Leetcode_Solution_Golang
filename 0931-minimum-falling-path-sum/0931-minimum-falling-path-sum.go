func minFallingPathSum(matrix [][]int) int {
    n := len(matrix)
    
    for i:=1; i<n; i++ {
        for j := range matrix[i] {
            matrix[i][j] = matrix[i][j]+min(matrix[i-1][max(j-1,0)], matrix[i-1][j], matrix[i-1][min(j+1, n-1)])
        }
    }
    
    minVal := math.MaxInt32
    for j := range matrix[0] {
        minVal = min(minVal, matrix[n-1][j])
    }
    
    return minVal
}

func min(vals ...int) int {
    minVal := math.MaxInt32
    
    for _, val := range vals {
        if val < minVal {
            minVal = val
        }
    }
    
    return minVal
}

func max(vals ...int) int {
    maxVal := math.MinInt32
    
    for _, val := range vals {
        if val > maxVal {
            maxVal = val
        }
    }
    
    return maxVal
}