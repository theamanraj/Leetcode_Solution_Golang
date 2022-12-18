func fillCups(amount []int) int {
    sort.Ints(amount)
    a, b := amount[0] + amount[1], amount[2]
    if a < b {
        return b
    }
    return b + (a - b + 1) / 2
}