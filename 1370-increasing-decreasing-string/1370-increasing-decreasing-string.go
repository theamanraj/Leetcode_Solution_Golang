func sortString(s string) string {
    counts := make([]int, 26)
    n := len(s)
    res := ""
    for _, c := range s {
        counts[c - 'a']++
    }
    for n > 0 {
        for i := 0; i < 26; i++ {
            if counts[i] > 0 {
                res += string(i + 'a')
                counts[i]--
                n--
            }
        }
        for i := 25; i >= 0; i-- {
            if counts[i] > 0 {
                res += string(i + 'a')
                counts[i]--
                n--
            }
        }
    }
    return res
}