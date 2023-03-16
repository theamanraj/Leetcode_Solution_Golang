func smallestStringWithSwaps(s string, pairs [][]int) string {
    graph := map[int][]int{}
    for _, pair := range pairs {
        a, b := pair[0], pair[1]
        graph[a] = append(graph[a], b)
        graph[b] = append(graph[b], a)
    }
    
    visited := map[int]bool{}
    answer := make([]byte, len(s))
    for v := range s {
        if !visited[v] {
            chars, indeces := []byte{}, []int{}
            visit(v,s, graph, visited, &chars, &indeces)
            sort.Slice(chars, func(x,y int) bool {
                return chars[x] < chars[y]
            })
            sort.Ints(indeces)
            
            for i:=0; i< len(chars);i++ {
                answer[indeces[i]] = chars[i]
            }
        }
        
    }
    return string(answer)
}

func visit(v int,s string, graph map[int][]int,visited map[int]bool,chars *[]byte, indeces *[]int) {
    *chars = append(*chars, s[v])
    *indeces = append(*indeces, v)
    
    visited[v] = true
    for _, n := range graph[v] {
        if !visited[n] {
            visit(n, s, graph, visited, chars, indeces)
        }
    }
}