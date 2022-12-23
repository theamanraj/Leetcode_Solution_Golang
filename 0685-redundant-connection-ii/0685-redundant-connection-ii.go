func find(hay []int, needle int) int {
    if hay[needle] != needle {
        return find(hay, hay[needle])
    }
    
    return needle
}

func union(hay []int, x, y int) {
    hay[x] = y;
}

func findRedundantDirectedConnection(edges [][]int) []int {
    sets := make([]int, len(edges) + 1)
    for i := range sets {
        sets[i] = i
    }
        
    parents := make([]int, len(edges) + 1)
    
    
    ans1, ans2 := []int{}, []int{}
    
    for _, edge := range edges {
        u, v := edge[0], edge[1]
        if parents[v] > 0 {
            ans1 = []int{parents[v], v}
            ans2 = edge
        }
        
        parents[v] = u
    }
    
    
    for _, edge := range edges {
        
        if len(ans2) == 2 && ans2[0] == edge[0] && ans2[1] == edge[1]{
            continue
        }
        
        x, y := find(sets, edge[0]), find(sets, edge[1])
        
        if x == y {
            if len(ans1) == 0  {
                return edge
            }
            return ans1
        }
        
        union(sets, x, y)
        
    }
       
    return ans2
}