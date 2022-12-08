package main

func minReorder(n int, connections [][]int) int {
	in, out := make(map[int][]int,n), make(map[int][]int,n)
	for i := 0; i < len(connections); i++ {
		in[connections[i][1]] = append(in[connections[i][1]], connections[i][0])
		out[connections[i][0]] = append(out[connections[i][0]], connections[i][1])
	}
	visited := make(map[int]bool, n)
	res := 0
	var dfs func(this int)
	dfs = func(this int) {
		visited[this] = true
		for i := 0; i < len(in[this]); i++ {
			if !visited[in[this][i]] {
				dfs(in[this][i])
			}
		}
		for i := 0; i < len(out[this]); i++ {
			if !visited[out[this][i]] {
				dfs(out[this][i])
				res++
			}
		}
	}
	dfs(0)
	return res
}