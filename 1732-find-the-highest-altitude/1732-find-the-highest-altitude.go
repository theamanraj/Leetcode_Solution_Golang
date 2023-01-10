func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
func largestAltitude(gain []int) int {
    var ans int = 0
    var curr int = 0
    for i := 0; i < len(gain); i++ {
        curr += gain[i]
        ans = max(ans, curr)
    }
    return ans
}