func reverse(s []byte) []byte {
    for i,j := 0,len(s)-1; i < j; i,j = i+1,j-1 {
        s[i], s[j] = s[j], s[i]
    }
    return s
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func addStrings(num1 string, num2 string) string {
    b1 := reverse([]byte(num1))
    b2 := reverse([]byte(num2))
    
    var carry byte = 0
    l := max(len(b1), len(b2))
    res := make([]byte, l)
    
    for i := 0; i < l; i++ {
        if i < len(b1) {
            carry += b1[i] - '0'
        }
        if i < len(b2) {
            carry += b2[i] - '0'
        }
        res[i] = carry % 10 + '0'
        carry /= 10
    }
    
    if carry != 0 {
        res = append(res, carry + '0')
    }
    return string(reverse(res))
}