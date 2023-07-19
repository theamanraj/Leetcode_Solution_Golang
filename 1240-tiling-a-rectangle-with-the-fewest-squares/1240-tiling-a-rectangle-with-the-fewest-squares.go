func tilingRectangle(m int, n int) int {

    mx, res := makeMx(m, n), m * n
    
    var dfs func(k, i, c int)
    dfs = func(k, i, c int) {
        
        switch {
        case c >= res: // already have better result
            return
        case k >= m: // all rows processed
            res = c
            return
        case i >= n: // all columns processed in current row
            dfs(k + 1, 0, c)
            return
        case mx[k][i]: // cell already occupied
            dfs(k, i + 1, c)
            return
        }

        for w := min(m - k, n - i); w >= 1; w-- {
            if canFill(k, i, w, mx) {            
                fillMx(k, i, w, true, mx)
                dfs(k, i + w, c + 1)
                fillMx(k, i, w, false, mx)
            }
        }
    }
    
    dfs(0, 0, 0)

    return res
}

func makeMx(m, n int) [][]bool {
    mx := make([][]bool, m)
    for k := range mx {
        mx[k] = make([]bool, n)
    }
    return mx
}

func canFill(k0, i0, w int, mx [][]bool) bool {
    for k, ck := k0, k0 + w; k < ck; k++ {
        for i, ci := i0, i0 + w; i < ci; i++ {
            if mx[k][i] {
                return false
            }
        }
    }
    return true
}

func fillMx(k0, i0, w int, x bool, mx [][]bool) {
    for k, ck := k0, k0 + w; k < ck; k++ {
        for i, ci := i0, i0 + w; i < ci; i++ {
            mx[k][i] = x
        }
    }
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}