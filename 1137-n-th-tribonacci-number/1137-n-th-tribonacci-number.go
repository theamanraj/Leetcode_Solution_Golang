func tribonacci(n int) int {
    cache := map[int]int{
        0 : 0, 1 : 1, 2 : 1,
    }
    return recursive(cache, n)
}

func recursive(cache map[int]int, n int) int {
    if _, ok := cache[n]; !ok {
        cache[n] = recursive(cache, n-3) + recursive(cache, n-2) + recursive(cache, n-1)
    }
    return cache[n]
}