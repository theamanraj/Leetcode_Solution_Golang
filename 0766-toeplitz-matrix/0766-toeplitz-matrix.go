func isToeplitzMatrix(matrix [][]int) bool {
    
    
    if len(matrix) == 0 {
        return true
    }
    
    for i := 0; i < len(matrix); i++ {
        t := Check(i, 0, matrix)
        if !t {
            return false
        }
    }
    
    for i := 0; i < len(matrix[0]); i++ {
        t := Check(0, i, matrix)
        if t == false {
            return false
        }
    }

    return true
}

func Check(n, m int, matrix [][]int) bool {
    
    k := matrix[n][m]
    
    for n < len(matrix) && m < len(matrix[0]) {
        if matrix[n][m] != k {
            return false
        }
        n += 1
        m += 1
    }
    return true
}