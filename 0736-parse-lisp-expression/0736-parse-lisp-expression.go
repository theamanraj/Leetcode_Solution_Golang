var variables map[string][]int

func isDigit(b byte) bool {
    return b == '-' || (0 <= b - '0' && b - '0' < 10)
}

func isLower(b byte) bool {
    return 0 <= b - 'a' && b - 'a' < 26
}

func evalNumber(s string) (int, int) {
    n, i := 0, 0
    negative := false
    if s[i] == '-' {
        negative = true
        i++
    }
    for ; i < len(s) && isDigit(s[i]); i++ {
        n = n * 10 + int(s[i]-'0')
    }
    if negative {
        n = -n
    }
    return n, i
}

func evalVariable(s string) (string, int, int) {
    i := 0
    for ; i < len(s) && (isDigit(s[i]) ||isLower(s[i])); i++ {}
    var n int
    v := s[:i]
    if _, exist := variables[v]; exist {
        n = variables[v][len(variables[v])-1]
    }
    
    return v, n, i
}

func evalLet(s string) (int, int) {
    local := make(map[string]bool)
    for i := 0; i < len(s); i++ {
        v, n, j := evalExpr(s[i:])
        i += j
        if s[i] == ')' {
            for v, _ := range local {
                variables[v] = variables[v][:len(variables[v])-1]
            }
            return n, i
        }
        i++
        _, n, j = evalExpr(s[i:])
        i += j
        if local[v] {
            variables[v][len(variables[v])-1] = n
        } else {
            local[v] = true
            variables[v] = append(variables[v], n)
        }
    }
    
    return 0, 0
}

func evalOp(s string) (int, int) {
    var n, i int
    if s[0] == 'l' {
        n, i = evalLet(s[4:])
        i += 4
    } else {
        i = 4
        if s[0] == 'm' {
            i = 5
        }
        _, first, j := evalExpr(s[i:])
        i += j + 1
        _, second, j := evalExpr(s[i:])
        i += j
        if s[0] == 'a' {
            n = first + second
        } else {
            n = first * second
        }
    }
    return n, i
}

func evalExpr(s string) (string, int, int) {
    if isDigit(s[0]) {
        n, i := evalNumber(s)
        return "", n, i
    } else if s[0] == '(' {
        n, i := evalOp(s[1:])
        return "", n, i + 2
    } else {
        return evalVariable(s)
    }
}

func evaluate(s string) int {
    variables = make(map[string][]int)
    _, n, _ := evalExpr(s)
    return n
}