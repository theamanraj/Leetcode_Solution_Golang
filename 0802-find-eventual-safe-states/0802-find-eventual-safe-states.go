const (
	Unknow   = iota
	Unsafe   = iota
	Safe     = iota
	Visiting = iota
)

func eventualSafeNodes(graph [][]int) []int {
	v := len(graph)
	marked := make([]int, v)
	for i := 0; i < v; i++ {
		dfs(graph, i, marked)
	}
	ans := make([]int, 0)
	for i := 0; i < v; i++ {
		if marked[i] == Safe {
			ans = append(ans, i)
		}
	}
	return ans
}

func dfs(graph [][]int, v int, marked []int) int {
	if marked[v] != Unknow {
		if marked[v] == Visiting {
			marked[v] = Unsafe
		}
		return marked[v]
	}
	marked[v] = Visiting

	for _, u := range graph[v] {
		if dfs(graph, u, marked) == Unsafe {
			marked[v] = Unsafe
			return marked[v]
		}
	}
	marked[v] = Safe
	return marked[v]
}