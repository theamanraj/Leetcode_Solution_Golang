func countLargestGroup(n int) int {
    groups := map[int]int{}
    for num := 1; num <= n; num++ {
        digitSum := 0
        for i := num; i > 0; i /= 10 {
            digitSum += i % 10
        }
        groups[digitSum]++
    }
    max := 0
    res := 0
    for _, v := range groups {
        if v > max {
            max = v
            res = 1
        } else if v == max {
            res++
        }
    }
    return res
}