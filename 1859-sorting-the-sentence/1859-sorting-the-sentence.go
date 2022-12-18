func sortSentence(s string) string {
    tmp := strings.Split(s, " ")
    ans := make([]string, len(tmp))
    for _, v := range tmp { ans[v[len(v) - 1] - '1'] = v[:len(v) - 1] }
    return strings.Join(ans, " ")
}