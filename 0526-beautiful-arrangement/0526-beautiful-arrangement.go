func helper(arr []int, idx, n, ans int) int {
    for i, v := range arr {
        if v % idx == 0 || idx % v == 0 {
            if idx == n {
                return ans + 1
            }
            narr := make([]int, len(arr))
            copy(narr, arr)
            narr = append(narr[:i], narr[i+1:]...)
            ans = helper(narr, idx +1, n, ans)
        }
    }
    
    return ans
}

func countArrangement(n int) int {
    arr := make([]int, n)
    for i := 1; i <= n; i++ {
        arr[i - 1] = i
    }
    return helper(arr, 1, n, 0)
}