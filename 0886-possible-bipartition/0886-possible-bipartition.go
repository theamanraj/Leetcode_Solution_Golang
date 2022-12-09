func possibleBipartition(n int, dislikes [][]int) bool {
    
    var color = make([]int, n+1)    
    var graph = make(map[int][]int)
    
    for _, v := range dislikes{
        graph[v[0]] = append(graph[v[0]], v[1])
        graph[v[1]] = append(graph[v[1]], v[0])
    }
    
    for k, _ := range graph{
		if color[k]==0 && !bfs(k,graph, &color){
			return false
		}
	}
	return true
}


func bfs(curr int, graph map[int][]int, color *[]int) bool{

	var q []int
	q = append(q, curr)
	(*color)[curr] = 1

	for len(q)>0{
		sizeLength := len(q)
		for i:=0; i<sizeLength; i++{

			currNode := q[0]
			q = q[1:]
			currColor := (*color)[currNode]
            
			for _, n := range graph[currNode]{

				//if neighbour is of same color/type
				if (*color)[n] == currColor{
					return false
				}

				//if not visited
				if (*color)[n] == 0{
					(*color)[n] = -1 * currColor
					q = append(q, n)
				}
			}
		}
	}
	return true
}