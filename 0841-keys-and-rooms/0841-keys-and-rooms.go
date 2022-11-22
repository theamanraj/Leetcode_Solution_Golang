func canVisitAllRooms(rooms [][]int) bool {
    /*
    graph := make([][]int, len(rooms))
    for i:=0; i < len(rooms); i++ {
        graph[i] = append(graph[i], rooms[i]...)
    }
    */
    visited := make([]bool, len(rooms))
    dfs(rooms, visited, 0)
    for _, val := range visited {
        if val == false { return false }
    }
    return true
}

func dfs(graph [][]int, visited []bool, node int) {
    if visited[node] == true { return }
    visited[node] = true
    neighbors := graph[node]
    for _, each := range neighbors {
        dfs(graph, visited, each)
    }
}