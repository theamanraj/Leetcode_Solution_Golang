func diffWaysToCompute(input string) []int {
    ans := []int{}
    for i := 0; i < len(input); i++ {
        if input[i] == '+' || input[i] == '-' || input[i] == '*' {
            for _, a := range diffWaysToCompute(input[:i]) {
                for _, b := range diffWaysToCompute(input[i + 1:]) {
                    switch input[i] {
                    case '+':
                        ans = append(ans, a + b)
                    case '-':
                        ans = append(ans, a - b)
                    case '*':
                        ans = append(ans, a * b)
                    }
                }
            }
        }
    }
    if tmp, _ := strconv.Atoi(input); len(ans) == 0 { ans = append(ans, tmp) }
    return ans
}