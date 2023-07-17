func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}

func countSquares(matrix [][]int) int {
    
    count := 0
    for r := 0; r < len(matrix); r++ {
        for c := 0; c < len(matrix[0]); c++ {
            if matrix[r][c] == 1 {
                if r == 0 || c == 0  {
                    count++
                } else {
                    matrix[r][c] = 1 + min(matrix[r-1][c-1], min(matrix[r-1][c], matrix[r][c-1]))
                    count += matrix[r][c]
                }
            }
        }
    }
    
    return count
}