func fib(n int) int {
    cur, next := 0, 1
    for i := 0; i < n; i++ {
        cur, next = next, cur + next
    }
    return cur
}