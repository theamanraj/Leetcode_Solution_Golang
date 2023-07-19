func circularPermutation(n int, start int) []int {
    size := 1 << n
    cp := make([]int, size)
    for i := 0; i < size; i++ {
        cp[i] =  i ^ i >> 1 ^ start
    }
    return cp
}