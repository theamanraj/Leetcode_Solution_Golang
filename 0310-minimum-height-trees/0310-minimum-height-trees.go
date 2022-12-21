func findMinHeightTrees(n int, edges [][]int) []int {
    if n == 0 {
        return []int{}
    }
    if n == 1 {
        return []int{0}
    }
    degrees := make([]int, n)
    adjacency := make([][]int, n)
    for _, edge := range edges {
        degrees[edge[0]]++
        degrees[edge[1]]++
        adjacency[edge[0]] = append(adjacency[edge[0]], edge[1])
        adjacency[edge[1]] = append(adjacency[edge[1]], edge[0])
    }
    
    q := []int{}
    for i := range degrees {
        if degrees[i] == 1 {
            q = append(q, i)
        }
    }
    
    for n > 2 {
        size := len(q)
        n -= size
        for size > 0 {
            v := q[0]
            q = q[1:]
            for _, v2 := range adjacency[v] {
                degrees[v2]--
                if degrees[v2] == 1 {
                    q = append(q, v2)
                }
            }
            size--
        }
    }
    
    return q
}