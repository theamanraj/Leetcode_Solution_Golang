func isSubsequence(s string, t string) bool {
    x,y := 0,0
    for x != len(s) && y != len(t){
        if s[x] == t[y]{
            x++
            y++
        } else {
            y++
        }
    }
    return x == len(s)
}