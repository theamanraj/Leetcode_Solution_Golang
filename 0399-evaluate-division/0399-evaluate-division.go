type Node struct{
    v string
    w float64
}


func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
    
    var graph = make(map[string][]Node)
    
    for i, e := range equations{
        graph[e[0]] = append(graph[e[0]], Node{v: e[1], w: values[i]}) 
        graph[e[1]] = append(graph[e[1]], Node{v: e[0], w: 1/values[i]}) 
    }
    
    var result = make([]float64, len(queries))
    for i, q := range queries{
        var visited = make(map[string]bool)
        result[i] = dfs(q[0], q[1], graph, visited)
    }
    
    return result
}


func dfs(src string, dst string, graph map[string][]Node, visited map[string]bool) float64{
    
    if _, ok := graph[src]; !ok{
        return -1.0
    }
    
    if src==dst{
        return 1.0
    }
    
    visited[src] = true
    
    for _, n := range graph[src]{
        if !visited[n.v]{
            ans := dfs(n.v, dst, graph, visited)
            if ans != -1.0{
                return ans * n.w
            }
        }
    }
    return -1.0
}
