type UnionFind struct{
    parents []int
}

func NewUnionFind(n int) UnionFind{
    parents := make([]int,n)
    
    for i:=0;i<n;i++{
        parents[i] = i
    }
    
    return UnionFind{
        parents : parents,
    }
}

func (uf *UnionFind) Find(i int) int{
    for i != uf.parents[i]{
        uf.parents[i] = uf.parents[uf.parents[i]]
        i = uf.parents[i]    
    }
    
    return uf.parents[i]
}

func (uf *UnionFind) Union(i int,j int) bool{
    parentP := uf.Find(i)
    parentQ := uf.Find(j)
    
    if parentP == parentQ{
        return false
    }
    
    if parentQ < parentP{
        uf.parents[parentP] = parentQ
        return true
    }
    
    uf.parents[parentQ] = parentP
    
    return true
}

// An IntHeap is a max-heap of ints.
type item struct{
    val int
    origin int
    destination int
}

type PriorityQueue []item

func (h PriorityQueue) Len() int{ return len(h) }
func (h PriorityQueue) Less(i, j int) bool { return h[i].val < h[j].val }
func (h PriorityQueue) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *PriorityQueue) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(item))
}
func (h *PriorityQueue) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}


func minCostConnectPoints(points [][]int) int {
    n := len(points)
    uf := NewUnionFind(n)
    h := PriorityQueue{}
    heap.Init(&h)
    
    for i:=0;i<n;i++{
        for j:=i+1;j<n;j++{
            cost := manhatDist(points[i],points[j])
            currEdge := item{
                val : cost,
                origin : i,
                destination : j,
            }
            heap.Push(&h,currEdge)
        }
    }
    
    minCost := 0
    for h.Len() > 0{
        currEdge := heap.Pop(&h).(item)
        isTree := uf.Union(currEdge.origin,currEdge.destination)
        // if the edge formed cycle, skip it
        if !isTree{
            continue
        }
        
        minCost += currEdge.val
    }
    return minCost
}

func min(a int,b int) int{
    if a < b{
        return a
    }
    
    return b
}

func manhatDist(v1 []int,v2 []int) int{
    return abs(v1[0]-v2[0])+abs(v1[1]-v2[1])
}

func abs(a int) int{
    if a < 0{
        return -a
    }
    
    return a
}