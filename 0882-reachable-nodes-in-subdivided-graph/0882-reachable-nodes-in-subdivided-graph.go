type Edge struct {
	To           int
	TotalNodesBetween int
	UnvisitedNodesBetween *int
}

type Marker struct {
	Pos       int
	MovesLeft int
}

type PriorityQueue []Marker

func (p PriorityQueue) Len() int {
	return len(p)
}

func (p PriorityQueue) Less(i, j int) bool {
	return p[i].MovesLeft > p[j].MovesLeft // prioritise markers with more moves left
}

func (p PriorityQueue) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *PriorityQueue) Push(x interface{}) {
	*p = append(*p, x.(Marker))
}

func (p *PriorityQueue) Pop() interface{} {
	h := *p
	n := len(h)
	*p = h[:n-1]
	return h[n-1]
}

func reachableNodes(edges [][]int, maxMoves int, n int) int {
	// using the idea of BFS, when a node has been reached, no other attempts to it should be tried.
	// why? Every other that comes after is either equal or less in terms of moves left
	graph := make(map[int][]*Edge)
	for _, edge := range edges {
		// from edge[0] -> edge[1]
		dist := edge[2]
		toEdge1 := &Edge{
			To:           edge[1],
			TotalNodesBetween: dist,
			UnvisitedNodesBetween: &dist,
		}
		graph[edge[0]] = append(graph[edge[0]], toEdge1)

		// from edge[1] -> edge[0]
		toEdge0 := &Edge{
			To:           edge[0],
			TotalNodesBetween: dist,
			UnvisitedNodesBetween: &dist,
		}
		graph[edge[1]] = append(graph[edge[1]], toEdge0)
	}

	pq := &PriorityQueue{
		Marker{
			Pos:       0,
			MovesLeft: maxMoves,
		},
	}
	heap.Init(pq)
	visited := make([]bool, n)
	res := 0
	for pq.Len() > 0 {
		cur := heap.Pop(pq).(Marker)
        if visited[cur.Pos] {
			continue
		}
		neighbours := graph[cur.Pos]
		res++  // count this node
		visited[cur.Pos] = true
		for _, neighbour := range neighbours {
			// Can make it across and no one attempted 
			if cur.MovesLeft > neighbour.TotalNodesBetween { 
				heap.Push(pq, Marker{
					Pos:       neighbour.To,
					MovesLeft: cur.MovesLeft - neighbour.TotalNodesBetween-1, // -1 to make it fully across
				})
				res += *neighbour.UnvisitedNodesBetween
				*neighbour.UnvisitedNodesBetween = 0 // To prevent double counting same nodes from another direction, set to zero
			} else { // partial completion, note this is safe meaning only 1 direction will be attempted at most once
                moved := min(cur.MovesLeft, *neighbour.UnvisitedNodesBetween)
				res += moved
				*neighbour.UnvisitedNodesBetween -= moved
			}
		}
	}
	return res
}

func min(a, b int) int {
    if a < b {
        return a
    }
    
    return b
}