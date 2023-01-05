func makesquare(matchsticks []int) bool {
    sum := 0
    for _, v := range matchsticks { sum += v }
    if sum % 4 != 0 { return false }
    sideLen, sides := sum / 4, make([]int, 4)
    sort.Slice(matchsticks, func(i, j int) bool { return matchsticks[i] > matchsticks[j] })
    var backtrack func(i int) bool 
    backtrack = func(i int) bool {
        if i == len(matchsticks) { return true }
        for j := 0; j < 4; j++ {
            if sides[j] + matchsticks[i] <= sideLen {
                sides[j] += matchsticks[i]
                if backtrack(i + 1) { return true }
                sides[j] -= matchsticks[i]
            }
        }
        return false
    }
    return backtrack(0)
}