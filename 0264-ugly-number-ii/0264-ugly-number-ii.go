func nthUglyNumber(n int) int {
    arr := make([]int, n)
    arr[0] = 1
    a, b, c := 0, 0, 0
    f2, f3, f5 := 0, 0, 0
    for i := 1; i < len(arr); i++ {
        f2 = 2 * arr[a]
        f3 = 3 * arr[b]
        f5 = 5 * arr[c]
        min := Min(f2, f3, f5)
        arr[i] = min
        if min == f2 { a++ }
        if min == f3 { b++ }
        if min == f5 { c++ }
    }
    return arr[n - 1]
}

func Min(args ...int) int {
    min := math.MaxInt32
    for _, v := range args {
        if v < min { min = v }
    }
    return min
}