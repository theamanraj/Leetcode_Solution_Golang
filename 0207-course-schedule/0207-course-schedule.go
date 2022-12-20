func canFinish(numCourses int, prerequisites [][]int) bool {
    graph := map[int][]int{}
    for _, each := range prerequisites {
        n1, n2 := each[0], each[1]
        graph[n2] = append(graph[n2], n1)
    }
    visited := make([]int, numCourses)
    for node, _ := range graph {
        if visited[node] == 0 {
            if hasCycle(graph, visited, node) == true {
                return false
            }
        }
    }
    return true
}

func hasCycle(graph map[int][]int, visited []int, node int) bool {
    if visited[node] == 2 {
        return false
    }
    if visited[node] ==1 { return true }
    visited[node] = 1
    for _, neigh := range graph[node] {
        if hasCycle(graph, visited, neigh) == true {
            return true
        }
    }
    visited[node] = 2 //all nodes are covered from this node
    return false
}