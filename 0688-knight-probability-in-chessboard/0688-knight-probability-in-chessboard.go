var dir = [8][2]int{{2, -1}, {2, 1}, {-2, -1}, {-2, 1}, {1, 2}, {-1, 2}, {1, -2}, {-1, -2}}

func compute(N, K, r, c int, dp map[[3]int]float64) float64 {
    if r < 0 || c < 0 || r >= N || c >= N {
        return 0
    }
    if K == 0 {
        return 1
    }
    if val, ok := dp[[3]int{K, r, c}]; ok {
        return val
    }
    var out float64
    for _, v := range dir {
        out += compute(N, K - 1, r + v[0], c + v[1], dp) / 8.0
    }
    dp[[3]int{K, r, c}] = out
    return out
}

func knightProbability(N int, K int, r int, c int) float64 {
    dp := make(map[[3]int]float64, N * N * K)
    return compute(N, K, r, c, dp)
}