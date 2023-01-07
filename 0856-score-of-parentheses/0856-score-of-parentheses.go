func scoreOfParentheses(S string) int {
    s1 := make([]int, 0)
    res := 0
    tmp := 0
    sum := 0
    data := []byte(S)
    for _, v := range data {
        if v == '(' {
            s1 = append(s1, 0)
        } else {
            tmp = s1[len(s1) - 1]
            if tmp == 0 {
                s1 = s1[:len(s1) - 1]
                s1 = append(s1, 1)
            } else {
                sum = 0
                for ;s1[len(s1) - 1] != 0; {
                    sum += s1[len(s1) - 1]
                    s1 = s1[:len(s1) - 1]
                } 
                s1 = s1[:len(s1) - 1]
                s1 = append(s1, 2 * sum)
            }
        }
    }
    for i:=0; i<len(s1); i++ {
        res += s1[i]
    }
    return res
}