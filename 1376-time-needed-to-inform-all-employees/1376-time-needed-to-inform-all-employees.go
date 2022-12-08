func numOfMinutes(n int, headID int, manager []int, informTime []int) int {
    m := map[int][]int{}
    
    for i:=0;i<n;i++ {
        m[manager[i]] = append(m[manager[i]], i)
    }
    
    return dfs(m, informTime, n, headID)
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func dfs(m map[int][]int, informTime []int, n int, cur int) int {
    if cur >= n {
        return 0
    }
    
    ret := informTime[cur]
    ch := 0
    for _, e := range m[cur] {
        child := dfs(m, informTime, n, e)
        ch = max(ch, child)
    }
    
    return ret + ch
}