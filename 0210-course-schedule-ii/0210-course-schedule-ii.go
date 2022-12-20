func findOrder(numCourses int, prerequisites [][]int) []int {
    inDegreeList, adjList := getInDegreeAndAdjList(numCourses, prerequisites)
    queue := []int{}
    visited := 0
    res := []int{}
    for course, degree := range inDegreeList {
        if degree == 0 {
            queue = append(queue, course)
            res = append(res, course)
            visited++
        }
    }
    for len(queue) > 0 {
        src := queue[0]
        queue = queue[1:]
        for _, dst := range adjList[src] {
            inDegreeList[dst]--
            if inDegreeList[dst] == 0 {
                queue = append(queue, dst)
                res = append(res, dst)
                visited++
            }
        }
    }
    if len(res) == numCourses {
        return res
    }
    return []int{}
}

func getInDegreeAndAdjList(numCourses int, prerequisites [][]int) ([]int, [][]int) {
    adjList := make([][]int, numCourses)
    inDegreeList := make([]int, numCourses)
    for i := range adjList {
        adjList[i] = []int{}
    }
    for _, p := range prerequisites {
        src, dst := p[1], p[0]
        inDegreeList[dst] += 1
        adjList[src] = append(adjList[src], dst)
    }
    return inDegreeList, adjList
}