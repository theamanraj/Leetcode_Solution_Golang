func lengthLongestPath(input string) int {
    
    // level -> pathlen
    memo := make(map[int]int)
    res := dfs(strings.Split(input, "\n"), 0, 0, memo)
    return res
}

func dfs(paths []string, index int, pathlen int, memo map[int]int) int {
    if index >= len(paths) {
        return pathlen
    }
    curr := 1
    i := 0
    // '\n' '\t' is one byte
    for ; i < len(paths[index]) && paths[index][i] == '\t'; i++ {
        curr++
    }
    res := 0
    if strings.Contains(paths[index], ".") {
        res = memo[curr-1] + len(paths[index][i:])
    }
    
    // pathlen endswith '/'
    memo[curr] = memo[curr-1] + len(paths[index][i:]) + 1
    if res < pathlen {
        return dfs(paths, index+1, pathlen, memo)
    } else {
        return dfs(paths, index+1, res, memo)
    }
}