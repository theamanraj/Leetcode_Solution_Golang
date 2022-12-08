func allPathsSourceTarget(graph [][]int) [][]int {
    var ans [][]int
	path := []int{0}
    helper(graph, path, &ans) 
    return ans
}

func helper(graph [][]int, path []int, ans *[][]int) {
    pos := path[len(path)-1]
    target := len(graph)-1
    for _, child := range graph[pos] {
	    // A slice is passed by reference, not value, so we must make a copy of the slice for each child.
        cp := make([]int, len(path))
        copy(cp, path)
        cp = append(cp, child)
        if child == target {
            *ans = append(*ans, cp)
        } else {
            helper(graph, cp, ans)
        }
    }
    return
}