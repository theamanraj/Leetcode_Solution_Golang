func dfs(current, n int, result *[]int) {
    if current <= n {
        *result = append(*result, current)
        for i := 0; i <= 9; i++ {
            if current * 10 + i <= n {
                dfs(current * 10 + i, n, result)
            }
        }
    }
}

func lexicalOrder(n int) []int {
    result := []int{}
    
    for i := 1; i <= 9; i++ {
        dfs(i, n, &result)
    }
    
    return result
}