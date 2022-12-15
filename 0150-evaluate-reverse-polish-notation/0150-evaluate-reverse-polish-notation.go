import (
    "strconv"
)
func evalRPN(tokens []string) int {
    stack := []int{}
    
    funcMap := map[string]func(a, b int)int {
        "+": func(a, b int)int{return a + b},
        "-": func(a, b int)int{return a - b},
        "*": func(a, b int)int{return a * b},
        "/": func(a, b int)int{return a / b},
    }
    for i := 0; i < len(tokens); i++ {
        v, err := strconv.Atoi(tokens[i])
        if err == nil {
            stack = append(stack, v)
        } else {
            l := len(stack)
            a := stack[l - 2]
            b := stack[l - 1]
            stack = stack[:l - 2]
            stack = append(stack, funcMap[tokens[i]](a, b))
        }
    }
    
    return stack[0]
}