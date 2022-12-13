func spiralOrder(matrix [][]int) []int {
    direction := [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
    var res []int
    var curDir int
    var i, j int
    for i >= 0 && i < len(matrix) && j >= 0 && j < len(matrix[0]) && matrix[i][j] != 101{
        res = append(res, matrix[i][j])
        matrix[i][j] = 101
        iNext, jNext := i + direction[curDir][0], j + direction[curDir][1]
        if iNext >= 0 && iNext < len(matrix) && jNext >= 0 && jNext < len(matrix[0]) && matrix[iNext][jNext] != 101 {
            i, j = iNext, jNext
        } else {
            curDir = (curDir + 1) % 4
            i, j = i + direction[curDir][0], j + direction[curDir][1]
        }
    }
    return res
}