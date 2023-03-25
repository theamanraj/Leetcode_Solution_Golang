func countPairs(n int, edges [][]int) int64 {
    result := nc2(n)
    
    c := getCounts(n, edges)
    for _, k := range c {
        result -= nc2(k)
    }
    
    return result
} 

func nc2(n int) int64 {
    return int64((n*(n-1))>>1)
}

func getCounts(n int, edges [][]int) []int {
    graph := map[int][]int{}
    
    for i := 0; i < n; i++ {
        graph[i] = []int{}
    }
    for _, edge := range edges {
        graph[edge[0]] = append(graph[edge[0]], edge[1]) 
        graph[edge[1]] = append(graph[edge[1]], edge[0])  
    }
    
    x := map[int]bool{}
    p := 0
    
    results := []int{}
    for i := 0; i < n; i++ {
        if _, ok := x[i]; ok {
            continue
        }
        
        getCount(i, graph, x)
        
        results = append(results, len(x)-p)
        p = len(x)
    }

    return results
}

func getCount(n int, graph map[int][]int, x map[int]bool) {
    x[n] = true
    
    for _, adj := range graph[n] {
        if _, ok := x[adj]; ok {
            continue
        }
        getCount(adj, graph, x)
    }
}