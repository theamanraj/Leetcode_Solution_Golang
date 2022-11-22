func findJudge(N int, trust [][]int) int {
    if len(trust) == 0 && N == 1 { return 1 }
    if len(trust) == 0 { return -1 }
    people := make([]int, N + 1)
    for _, v := range trust {
        people[v[0]]--
        people[v[1]]++
    }
    for i, v := range people {
        if v == N - 1 { return i }
    }
    return -1
}