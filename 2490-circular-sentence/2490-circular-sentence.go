func isCircularSentence(sentence string) bool {
    s := strings.Split(sentence, " ")
    for i := 0; i < len(s); i++ {
        var s2 int
        if i == len(s) - 1 {
            s2 = 0
        } else {
            s2 = i + 1
        }
        if s[i][len(s[i])-1] != s[s2][0] {
            return false
        }
    }
    return true
}