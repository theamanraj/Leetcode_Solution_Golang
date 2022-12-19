func validPath(n int, edges [][]int, source int, destination int) bool {
    visited := make([]bool, n)
    al := make([][]int, n)
    for _, edge := range edges {
        al[edge[0]] = append(al[edge[0]], edge[1])
        al[edge[1]] = append(al[edge[1]], edge[0])
    }

    return pathExists(source, destination, al, visited)
}

func pathExists(s int, d int, al [][]int, v []bool) bool {
    if s == d {
        return true
    }

    if v[s] {
        return false
    }
    v[s] = true

    for _, child := range al[s] {
        if v[child] {
            continue
        }

        if pathExists(child, d, al, v) {
            return true
        }
    }
    return false
}