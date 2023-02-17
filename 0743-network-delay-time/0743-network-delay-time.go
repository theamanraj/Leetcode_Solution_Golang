func networkDelayTime(times [][]int, n int, k int) int {
    adj := make([][]pair, n + 1)    
    for i:=0; i < len(times); i++{
        e1, e2 := times[i][0], pair{times[i][1], times[i][2]}
        adj[e1] = append(adj[e1], e2)
    }    
    visited := make(map[int]int)
    minHeap := &MinHeap{}
    minHeap.Push(pair{k, 0})       
    ans := -1
    for len(*minHeap) > 0{        
        curr := heap.Pop(minHeap).(pair)  
        if visited[curr.node] > 0{
            continue
        }
        ans = curr.time
        visited[curr.node]++
        for _, v := range adj[curr.node]{
            if visited[v.node] == 0{
                heap.Push(minHeap, pair{v.node, v.time + curr.time})
            }
        }                    
    }
    if len(visited) < n{
        return -1
    }
    return ans
}

type pair struct{
    node int
    time int
}

type MinHeap []pair

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].time < h[j].time }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(pair))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}