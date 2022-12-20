func calculate(s string) int {
    var buf bytes.Buffer
    
    // Remove all white spaces
    for _, r := range s {
        if unicode.IsNumber(r) || r == '+' || r == '-' || r == '/' || r == '*' {
            buf.WriteRune(r)
        }
    }
    
    if buf.Len() == 0 {
        return 0
    }
    
    s = buf.String()
    l := len(s)
    i := 0
    
    
    
    for i < l && unicode.IsNumber(rune(s[i])) {i++}
    cur, _ := strconv.Atoi(s[:i])
    
    res := 0
    for i < l {
        oper := s[i]
        i++
        
        start := i
        for i < l && unicode.IsNumber(rune(s[i])) {i++}
        a, _ := strconv.Atoi(s[start:i])
        
        switch oper {
            case '+': 
            res += cur
            cur = a
            case '-': 
            res += cur
            cur = -1*a
            case '*':
            cur *= a
            case '/':
            cur /= a
        }
    }   
    
    return res+cur
}