func max(nums []int) int{
    tmp := nums[0]
    for _, n := range nums {
        if n > tmp {
            tmp = n
        }
    }
    return tmp
}
func DFS(arr []int, start, end int, mct *[][]int) int {
    if end - start == 1 {
        return 0
    }
    min := -1
    for i := start + 1; i < end; i ++{
        if (*mct)[start][i - 1] == 0 {
            (*mct)[start][i - 1] = DFS(arr, start, i, mct)
        }
        if (*mct)[i][end - 1] == 0 {
            (*mct)[i][end - 1] = DFS(arr, i, end, mct)
        }
        sum := max(arr[start : i]) * max(arr[i : end]) + (*mct)[start][i - 1] + (*mct)[i][end - 1]
        if sum < min || min == -1 {
            min = sum
        }
    }
    return min
}
func mctFromLeafValues(arr []int) int {
    mct := make([][]int, len(arr))
    for i := 0; i < len(arr); i ++ {
        mct[i] = make([]int, len(arr))
        mct[i][i] = 0
        if i != len(arr) - 1 {
            mct[i][i + 1] = arr[i] * arr[i + 1]
        }
    }
    return DFS(arr, 0, len(arr), &mct)
}