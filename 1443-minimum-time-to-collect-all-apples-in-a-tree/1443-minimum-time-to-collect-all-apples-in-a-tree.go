type MinTime struct {
	graph    map[int][]int
	hasApple []bool
	visited  []bool
}

func minTime(n int, edges [][]int, hasApple []bool) int {
	graph := map[int][]int{}
	for _, edge := range edges {
		n1, n2 := edge[0], edge[1]
		graph[n1] = append(graph[n1], n2)
		graph[n2] = append(graph[n2], n1)
	}
	visited := make([]bool, n)
	obj := &MinTime{graph, hasApple, visited}
	resp, exists := obj.Solve(0)
	if exists == true {
		return resp
	}
	return 0
}

func (obj *MinTime) Solve(start int) (int, bool) {
	t, e := 0, false
	obj.visited[start] = true
	for _, neighbor := range obj.graph[start] {
		if obj.visited[neighbor] == true {
			continue
		}
		t1, exists := obj.Solve(neighbor)
		if exists == true {
			t = t + t1 + 2
			e = e || exists
		}
	}
	if obj.hasApple[start] == true {
		return t, true
	}
	return t, e
}