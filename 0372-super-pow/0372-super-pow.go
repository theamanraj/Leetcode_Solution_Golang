func myPow(x int, y int) int {
    result := 1
    for i := 0; i < y; i++ {
        result = (result * x) % 1337
    }
    return result
}

func superPow(a int, b []int) int {
    a = a % 1337
    base := a
    result := myPow(base, b[len(b) - 1])
    for i := len(b) - 2; i > -1; i-- {
        base = myPow(base, 10)
        if base == 0 {
            return 0
        }
        result = (result * myPow(base, b[i])) % 1337
    }
    return result
}