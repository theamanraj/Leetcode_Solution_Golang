func spreadI(i, j, m, n int, a [][]byte) {
    if a[i][j] != 'O' {
        return
    }
    
    a[i][j] = 'I'
    
    if i < m-1 {
        spreadI(i+1, j, m, n, a)
    }
    
    if i > 0 {
        spreadI(i-1, j, m, n, a)
    }
    
    if j < n-1 {
        spreadI(i, j+1, m, n, a)
    }
    
    if j > 0 {
        spreadI(i, j-1, m, n, a)
    }
}


func solve(a [][]byte)  {
    m := len(a)
    
    if m == 0 {
        return
    }
    
    n := len(a[0])

    for j := 0; j < n; j++ {
        if a[0][j] == 'O' {
            spreadI(0, j, m, n, a)
        }
        if a[m-1][j] == 'O' {
            spreadI(m-1, j, m, n, a)
        }
    }
    
    for i := 0; i < m; i++ {
        if a[i][0] == 'O' {
            spreadI(i, 0, m, n, a)
        }
        if a[i][n-1] == 'O' {
            spreadI(i, n-1, m, n, a)
        }
    }
    
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if a[i][j] == 'O' {
                a[i][j] = 'X'
            }
        }
    }
    
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if a[i][j] == 'I' {
                a[i][j] = 'O'
            }
        }
    } 
}