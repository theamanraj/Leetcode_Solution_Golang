func numDecodings(s string) int {
    v := make([]int, len(s) + 1)
    v[0] = 1
    for i := range s {
        if i == 0 {
            if s[i] == '*' {
                v[1] = 9
            } else if s[i] != '0' {
                v[1] = 1
            }
        } else {
            fmt.Println(s[i] == '*')
            if s[i] == '*' {
                v[i + 1] += 9 * v[i]
            } else if s[i] != '0' {
                v[i + 1] += v[i]
            }
            if (s[i - 1] == '1' && s[i] != '*' ) || (s[i - 1] == '2' && s[i] <= '6' && s[i] >= '0') {
                v[i + 1] += v[i - 1]
            } else if s[i - 1] == '1' && s[i] == '*' {
                v[i + 1] += 9 * v[i - 1]
            } else if s[i - 1] == '*' && s[i] >= '0' && s[i] <= '6' {
                v[i + 1] += 2 * v[i - 1]
            } else if s[i - 1] == '*' && s[i] > '6' {
                v[i + 1] += v[i - 1]
            } else if s[i - 1] == '2' && s[i] == '*' {
                v[i + 1] += 6 * v[i - 1]
            } else if s[i] == '*' && s[i - 1] == '*' {
                v[i + 1] += 15 * v[i - 1]
            }
            
        }
        v[i + 1] = v[i + 1] % 1000000007
    }
    
    return v[len(v) - 1]
}