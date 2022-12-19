func maxScoreSightseeingPair(values []int) int {
    curMax := values[0]
    max := values[0]
    for i:=1; i<len(values); i++ {
        curMax--
        sum := curMax + values[i]
        if sum > max {
            max = sum
        } 
        
        if values[i] > curMax {
            curMax = values[i]
        }
    }
    return max
}