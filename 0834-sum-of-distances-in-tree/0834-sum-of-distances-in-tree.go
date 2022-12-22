/* Post-order */
func dfs(node int, parent int, adjList map[int][]int, count []int, result []int) {
    
    for _, child := range adjList[node] {
        if child == parent {
            continue 
        }
        
        dfs(child, node, adjList, count, result)
        count[node] += count[child]
        result[node] += result[child] + count[child]
    }   
}

/* Pre-order */
func dfs2(node int, parent int, adjList map[int][]int, count []int, result []int) {
    
    for _, child := range adjList[node] {
        if child == parent {
            continue 
        }
        
        result[child] = result[node] - count[child] + len(count) - count[child]
        dfs2(child, node, adjList, count, result)
    }    
}

func sumOfDistancesInTree(N int, edges [][]int) []int {
    
    adjList := make(map[int][]int)
    
    for _, edge := range edges {
        adjList[edge[0]] = append(adjList[edge[0]], edge[1])
        adjList[edge[1]] = append(adjList[edge[1]], edge[0])
    }
        
    count := make([]int, N)
    for i := 0; i < len(count); i++ {
        count[i] = 1
    }
    
    result := make([]int, N)
    
    dfs(0, -1, adjList, count, result)
    dfs2(0, -1, adjList, count, result)
    
    return result
}