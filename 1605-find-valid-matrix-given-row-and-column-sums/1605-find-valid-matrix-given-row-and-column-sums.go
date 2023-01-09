func restoreMatrix(rowSum []int, colSum []int) [][]int {
    m := len(rowSum)
    n := len(colSum)
    res := make([][]int,m)
    for i:=0;i<m;i++{
        res[i] = make([]int,n)
    }
    
    
    for i:=0;i<m;i++{
        for j:=0;j<n;j++{
            try := min(rowSum[i], colSum[j])
            res[i][j] = try
            rowSum[i]-=try
            colSum[j]-=try
            //if 0 break this row
            if rowSum[i] == 0 {
                break
            }
        }
    }
    return res
}

func min(a,b int)int{
    if a < b{
        return a
    }
    return b
}