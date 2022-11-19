func minRemoveToMakeValid(s string) string {
    n := len(s)
    if n == 0 {
        return ""
    }
    stack := []int{}
    buffer := make([]rune, n)
    for i, v := range s {
        if v == '(' {
            stack = append(stack, i)
        }
        if v == ')' {
            if len(stack) == 0 {
                buffer[i] = '#'
                continue
            }
            stack = stack[:len(stack)-1]
        }
        buffer[i] = v
    }
    for i := 0; i < len(stack); i++ {
        buffer[stack[i]] = '#'
    }
    var sb strings.Builder
    for _, v := range buffer {
        if v != '#' {
            sb.WriteRune(v)
        }
    }
    return sb.String()
    return ""
}