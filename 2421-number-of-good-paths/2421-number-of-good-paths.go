type UnionFind struct {
    root []int
}

func NewUF(n int) *UnionFind {
    root := make([]int, n)
    for i := 0; i < n; i++ {
        root[i] = i
    }
    return &UnionFind{root}
}

func (u *UnionFind) Find(x int) int {
    if u.root[x] != x {
        u.root[x] = u.Find(u.root[x])
    }
    return u.root[x]
}

func (u *UnionFind) Union(x, y int) {
    x, y = u.Find(x), u.Find(y)
    u.root[x] = y
}

func numberOfGoodPaths(vals []int, edges [][]int) int {
    n := len(vals)
    
    val2Node := make(map[int][]int, n)
    for i, val := range vals {
        val2Node[val] = append(val2Node[val], i)
    }
    
    val2Edge := make(map[int][][]int, n)
    for _, edge := range edges {
        j, k := edge[0], edge[1]
        maxVal := max(vals[j], vals[k])
        val2Edge[maxVal] = append(val2Edge[maxVal], []int{j, k})
    }
    
    uf := NewUF(n)
    res := 0
    
    sortedUniqueVals := make([]int, 0 , len(val2Node))
    for k, _ := range val2Node {
        sortedUniqueVals = append(sortedUniqueVals, k)
    }
    
    sort.Slice(sortedUniqueVals, func(i, j int) bool {
        return sortedUniqueVals[i] < sortedUniqueVals[j]
    })
    
    for _, val := range sortedUniqueVals {
        for _, edge := range val2Edge[val] {
            uf.Union(edge[0], edge[1])
        }
        count := make(map[int]int)
        for _, node := range val2Node[val] {
            count[uf.Find(node)] ++
        }
        for _, groupCount := range count {
            res += int((groupCount * (groupCount + 1)) / 2)
        }
    } 
    
    
    return res
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}