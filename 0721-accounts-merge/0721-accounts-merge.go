func accountsMerge(accounts [][]string) [][]string {
    edges := make(map[int][]int)
    emailToIdx := make(map[string]int)
    for i, acc := range accounts {
        for j := 1; j < len(acc); j++ {
            if idx, exists := emailToIdx[acc[j]]; exists {
                edges[idx] = append(edges[idx], i)
                edges[i] = append(edges[i], idx)
            } else {
                emailToIdx[acc[j]] = i
            }
        }
    }
    
    var res [][]string
    visited := make(map[int]bool)
    for i, acc := range accounts {
        if visited[i] {
            continue
        }
        // bfs
        queue := []int{i}
        acctEmails := make(map[string]bool)
        for len(queue) > 0 {
            idx := queue[0]
            queue = queue[1:]
            if !visited[idx] {
                visited[idx] = true
                for _, email := range accounts[idx][1:] {
                    acctEmails[email] = true
                }
                for _, neighbor := range edges[idx] {
                    queue = append(queue, neighbor)
                }
            }
        }
        mergedAcc := []string{acc[0]}
        for email := range acctEmails {
            mergedAcc = append(mergedAcc, email)
        }
        sort.Strings(mergedAcc[1:])
        res = append(res, mergedAcc)
    }
    return res
}