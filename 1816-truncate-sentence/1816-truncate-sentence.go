func truncateSentence(s string, k int) string {
    ans := []byte{}
    for _, c := range []byte(s) {
        if c == ' ' { k-- }
        if k == 0 { break }
        ans = append(ans, c)
    }
    return string(ans)
}