func diagonalSum(mat [][]int) int {
    sum := 0
    for k:=0;k<len(mat);k++ {
        sum+=mat[k][k]
        if k != len(mat)-1-k {
            sum+=mat[k][len(mat)-1-k]
        }
    }
    return sum
}