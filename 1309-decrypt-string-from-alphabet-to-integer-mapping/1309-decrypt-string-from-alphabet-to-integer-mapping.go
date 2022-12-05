func freqAlphabets(s string) string {
    res := []byte{}
    i := len(s) - 1
    for i >= 0 {
        if s[i] == '#' {
            n, _ := strconv.Atoi(s[i - 2:i])
            res = append([]byte{byte(n) - 1 + 'a'}, res...)
            i -= 3
        } else {
            n := s[i] - '0'
            res = append([]byte{n - 1 + 'a'}, res...)
            i--
        }
    }
    return string(res)
}