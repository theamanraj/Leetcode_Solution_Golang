func shortestPathLength(graph [][]int) int {
    visited := make([]map[int]bool, len(graph))
    
    queue := make([][]int, 0)
    for i := range graph {
        mask := 1 << i
        visited[i] = map[int]bool{mask: true}
        queue = append(queue, []int{i, mask})
    }
    
    target := 0
    for i := range graph {
        target = target | (1 << i)
    }
    
    res := 0
    for len(queue) > 0 {
        size := len(queue)
        for i := 0; i<size; i++ {
            cur := queue[0]
            queue = queue[1:]

            if cur[1] == target {
                return res
            }
            
            for _, neigh := range graph[cur[0]] {
                nextMask := cur[1] | (1 << neigh)
                if !visited[neigh][nextMask] {
                    if nextMask == target { // minor optimization to skip 1 level, not neccessary
                        return res + 1
                    }
                    visited[neigh][nextMask] = true
                    queue = append(queue, []int{neigh, nextMask})
                }
            }
        }
        
        res++
    }
    
    
    return res
}