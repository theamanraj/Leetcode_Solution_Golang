func isBipartite(graph [][]int) bool {
    visited := make([]int, len(graph))
    for i := 0; i < len(graph); i++ {
        if visited[i] != 0 { continue }
        queue := []int{ i }
        visited[i] = 1
        for len(queue) > 0 {
            cur := queue[0]
            queue = queue[1:]
            curM := visited[cur]
            neighborM := 0
            if curM == 1 {
                neighborM = 2
            } else {
                neighborM = 1
            }
            for _, neighbor := range graph[cur] {
                if visited[neighbor] == 0 {
                    visited[neighbor] = neighborM
                    queue = append(queue, neighbor)
                } else if visited[neighbor] != neighborM {
                    return false
                }
            }
        }
    }
    return true
}