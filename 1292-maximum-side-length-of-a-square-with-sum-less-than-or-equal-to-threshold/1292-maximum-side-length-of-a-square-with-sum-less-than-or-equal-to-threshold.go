func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func sumSquareEndingAt(sumToTopLeft [][]int, row, col, length int) int {
    return sumToTopLeft[row][col] -
        sumToTopLeft[row - length][col] -
        sumToTopLeft[row][col - length] +
        sumToTopLeft[row - length][col - length]
}

func maxSideLength(mat [][]int, threshold int) int {
    sumToTopLeft := make([][]int, len(mat)+1)
    sumToTopLeft[0] = make([]int, len(mat[0])+1)
    maxLen := 0
    for r, row := range mat {
        sumToTopLeft[r+1] = make([]int, len(row)+1)
        for c, value := range row {
            sumToTopLeft[r+1][c+1] =
                sumToTopLeft[r+1][c] +
                sumToTopLeft[r][c+1] -
                sumToTopLeft[r][c] +
                value

            if maxLen <= min(r, c) && sumSquareEndingAt(sumToTopLeft, r + 1, c + 1, maxLen + 1) <= threshold {
                maxLen++
            }
        }
    }
    
    return maxLen
}