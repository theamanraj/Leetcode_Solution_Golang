type UnionFind struct {
    words map[string]string
    size map[string]int
    groups int
}

func NewUnionFind() *UnionFind {
    return &UnionFind{
        words: map[string]string{},
        size: map[string]int{},
    }
}

func (uf *UnionFind) Find(p string) string {
    if uf.words[p] == p {
        return p
    }
    
    uf.words[p] = uf.Find(uf.words[p])
    return uf.words[p]
}

func (uf *UnionFind) Add(p string) {
    if _, ok := uf.words[p]; !ok {
        uf.words[p] = p
        uf.size[p] = 1
        uf.groups++
    }
}

func (uf *UnionFind) Union(p, q string) {
    
    u := uf.Find(p)
    v := uf.Find(q)
    
    if v == u {
        return
    }
    
    if uf.size[u] < uf.size[v] {
        uf.words[u] = v
        uf.size[v] += uf.size[u]
    } else {
        uf.words[v] = u 
        uf.size[u] += uf.size[v]
    }
    uf.groups--
}



func numSimilarGroups(strs []string) int {
    
    uf := NewUnionFind()

    for i, s := range strs {
        uf.Add(s)        
        for j := i + 1; j < len(strs); j++ {
            u := strs[j]
            uf.Add(u)
            
            if similar(u, s) {
                uf.Union(u, s)
            }
        }
    }    
    
    return uf.groups
}

func similar(first string, second string) bool {
    if len(first) != len(second) {
        return false
    }

    if first == second {
        return true
    }
    
    diff := 0
    for i := range first {
        cf := first[i]
        cs := second[i]
        
        if cf != cs {
            diff++
            if diff > 2 {
                return false
            }
        }
    }
    
    return diff <= 2
}