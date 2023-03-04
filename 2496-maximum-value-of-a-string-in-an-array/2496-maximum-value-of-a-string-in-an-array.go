func maximumValue(strs []string) int {
    var ans int
    for _, str := range strs {
        v, err := strconv.Atoi(str)
        if err == nil {
            ans = max(ans, v)
        } else {
            ans = max(ans, len(str))
        }
    }
    return ans
}

func max(i, j int) int {
    if i > j {
        return i
    }
    return j
}