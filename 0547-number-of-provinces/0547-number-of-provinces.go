func findCircleNum(M [][]int) int {
    
    var ans int
    visit := make([]bool, len(M))
    for i := 0; i < len(M); i++ {
        if !visit[i] {
            visit[i] = true
            dfs(M, i, visit)
            ans++    
        }
    }
    return ans
}

func dfs(M [][]int, me int, visit []bool) {
    for i := 0; i < len(M); i++ {
        if !visit[i] && M[me][i] == 1 {
            visit[i] = true
            dfs(M, i, visit)
        }
    }
}