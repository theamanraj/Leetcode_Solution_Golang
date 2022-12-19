func decodeString(s string) string {
    result, _ := helper(s, 0)
    return result
}

func helper(s string, i int) (string, int) {
    count := 0
    chunk := ""
    number := ""
    tmp := ""
    for i < len(s) {
        if s[i] == '[' {
            count, _ = strconv.Atoi(number)
            i++
            number = ""
            tmp, i = helper(s, i)
            for k := 0; k < count; k++ {
                chunk += tmp
            }
        } else if s[i] == ']' {
            return chunk, i
        } else if s[i] >= '0' && s[i] <= '9' {
            number += string(s[i])
        } else {
            chunk += string(s[i])
        }
        i++
    }
    return chunk, i
}