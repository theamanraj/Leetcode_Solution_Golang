func valid(s string) bool{
    if len(s) == 1 || s[0] != '0' && s[len(s) - 1] != '0' || s[len(s) - 1] == '0' && s[0] != '0' || s[0] == '0' && s[len(s) - 1] != '0' {
        return true
    }
    return false
}
func addDot(s string) []string{
    if len(s) == 1 || s[len(s) - 1] == '0' {
        return []string{s}
    }
    if s[0] == '0' {
        return []string{"0." + s[1 :]}
    }    
    res := []string{s}
    for i := 1; i < len(s); i ++{
        res = append(res, s[0 : i] + "." + s[i : ])
    }
    return res
} 
func ambiguousCoordinates(S string) []string {
    var res []string
    S = S[1 : len(S) - 1]
    for i := 1; i < len(S); i ++{
        if valid(S[0 : i]) && valid(S[i : ]) {
            left := addDot(S[0 : i])
            right := addDot(S[i :])
            for _, l := range left {
                for _, r := range right {
                    res = append(res, "(" + l + ", " + r + ")")
                }
            }
        }
    }    
    return res
}