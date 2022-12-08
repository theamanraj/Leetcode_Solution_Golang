func makeConnected(n int, connections [][]int) int {
	var visited []bool = make([]bool,n)
	var graph map[int]map[int]bool = make(map[int]map[int]bool)
	var l int = len(connections)
	for i := 0;i < l;i++{ // Build relationship for node
		if _,ok := graph[connections[i][0]];!ok{
			graph[connections[i][0]] = make(map[int]bool)
		}
		graph[connections[i][0]][connections[i][1]] = true
		if _,ok := graph[connections[i][1]];!ok{
			graph[connections[i][1]] = make(map[int]bool)
		}
		graph[connections[i][1]][connections[i][0]] = true
	}
	var groups int = 0
	var dup_lines int = 0
	for i := 0;i < n;i++{
		if visited[i]{
			continue
		}
		dup_lines += dfs_makeConnected(-1,i,graph,visited)
		groups++
	}
	if dup_lines < groups - 1 {
		return -1
	}
	return groups - 1
}

func dfs_makeConnected(pre_node int,cur_node int,graph map[int]map[int]bool,visited []bool)int{
	if visited[cur_node]{
		return 0
	}
	visited[cur_node] = true
	if _,ok := graph[cur_node];!ok{
		return 0
	}
	var dup_lines int = 0
	for neigh,_:= range graph[cur_node]{
		if visited[neigh]{
			if pre_node != neigh && pre_node != -1{
				dup_lines++
			}
		}
		dup_lines += dfs_makeConnected(cur_node,neigh,graph,visited)
	}
	return dup_lines
}