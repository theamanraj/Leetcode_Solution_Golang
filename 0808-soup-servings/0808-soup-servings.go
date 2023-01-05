func soupServings(N int) float64 {
    if N > 1e4 {
        return 1
    }
    N = int(math.Ceil(float64(N) / 25))
    mem := make([][]float64, N + 1)
    for i := 0; i <= N; i++ {
        mem[i] = make([]float64, N + 1)
    }
    return dfs(N, N, mem)
}

func dfs(a, b int, mem [][]float64) float64 {
    if a <= 0 && b <= 0 {
        return 0.5
    }
    if a <= 0 {
        return 1
    }
    if b <= 0 {
        return 0
    }
    if mem[a][b] > 0 {
        return mem[a][b]
    }
    mem[a][b] = 0.25 * (dfs(a - 4, b, mem) + dfs(a - 3, b - 1, mem) + dfs(a - 2, b - 2, mem) + dfs(a - 1, b - 3, mem))
    return mem[a][b]
}