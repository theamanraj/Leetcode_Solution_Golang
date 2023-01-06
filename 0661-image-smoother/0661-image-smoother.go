func imageSmoother(img [][]int) [][]int {
    res := make([][]int, len(img))
    for i := 0; i < len(img); i++ {
        res[i] = make([]int, len(img[0]))
    }
    
    for i := 0; i < len(img); i++ {
        for j := 0; j < len(img[0]); j++ {
            res[i][j] = sumMatrix(img, i, j)
        }
    }
    
    return res
}

func sumMatrix(m [][]int, x int, y int) int {
    sum, num := m[x][y], 1
    
    directions := [][]int{
        {-1, -1}, {-1, 0}, {-1, 1},
        {0, -1}, {0, 1},
        {1, -1}, {1, 0}, {1, 1},
    }
    
    for _, d := range directions {
        v, n := matrixValue(m, x + d[0], y + d[1])
        sum += v
        num += n
    }
    
    return sum / num
}

func matrixValue(m [][]int, x int, y int) (int, int) {
    if x < 0 || x >= len(m) || y < 0 || y >= len(m[0]) {
        return 0, 0
    }
    
    return m[x][y], 1
}