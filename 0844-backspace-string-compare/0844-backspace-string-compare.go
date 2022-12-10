func backspaceCompare(s string, t string) bool {
    var s1,s2 []byte
    for i := 0;i < len(s); i++ {
        if s[i] == '#' {
            if len(s1) > 0 { s1 = s1[:len(s1)-1] }
            continue
        }
        s1 = append(s1, s[i])
    }
    for i := 0; i<len(t); i++ {
        if t[i] == '#' {
            if len(s2) > 0 { s2 = s2[:len(s2)-1] }
            continue
        }
        s2 = append(s2, t[i])
    }
    return string(s1) == string(s2)
}