func countNumbersWithUniqueDigits(n int) int {
    if n == 0 {
        return 1
    }
    res, temp, available := 10, 9, 9
    for n  > 1 && available > 0 {
        temp = temp * available
        res += temp
        n--
        available--
    }
    return res
}