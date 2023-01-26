type CityCost struct {
	name int
	cost int
	stop int
}

type minHeap []CityCost

func (m minHeap) Len() int {
	return len(m)
}

func (m minHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m minHeap) Less(i, j int) bool {
	return m[i].cost < m[j].cost
}

func (m *minHeap) Push(x interface{}) {
	*m = append(*m, x.(CityCost))
}

func (m *minHeap) Pop() interface{} {
	n := len(*m)
	p := (*m)[n-1]
	*m = (*m)[:n-1]
	return p
}

type City struct {
	name int
	cost int
}

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	visited := make([]int, n)
	for i := 0; i < n; i++ {
		visited[i] = -1
	}

	adj := make(map[int][]City)
	for _, flight := range flights {
		source := flight[0]
		dest := flight[1]
		cost := flight[2]

		adj[source] = append(adj[source], City{name: dest, cost: cost})
	}

	minHeap := &minHeap{}
	heap.Init(minHeap)

	minHeap.Push(CityCost{
		name: src,
		cost: 0,
		stop: k+1,
	})

	for len(*minHeap) != 0{
		currNode := heap.Pop(minHeap).(CityCost)
		name := currNode.name
		cost := currNode.cost
		stop := currNode.stop

		if name == dst {
			return cost
		}
		
		if stop == 0 {
			continue
		}

		if visited[name] != -1 && visited[name] >= stop {
			continue
		}
		visited[name] = stop

		for _, city := range adj[name]{
			nextCost := city.cost
			nextName := city.name

			heap.Push(minHeap, CityCost{
				name: nextName,
				cost: nextCost + cost,
				stop: stop-1,
			})
		}
	}

	return -1
}