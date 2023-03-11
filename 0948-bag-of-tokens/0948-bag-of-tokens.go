func bagOfTokensScore(tokens []int, power int) int {
    sort.Ints(tokens)
    
    result := 0
    for i, j := 0, len(tokens)-1; i <= j; {
        if tokens[i] <= power {
            power -= tokens[i]
            result++
            i++
        } else if result == 0 || i == j {
            break
        } else {
            result--
            power += tokens[j]
            j--
        }
    }
    
    return result
}