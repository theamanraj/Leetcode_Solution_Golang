func reformat(s string) string {
    charCount, digitCount := 0, 0
    chars, digits, res := []byte{}, []byte{}, []byte{}
    for _, v := range []byte(s) {
        if '0' <= v && v <= '9' {
            digitCount++
            digits = append(digits, v)
        } else {
            charCount++
            chars = append(chars, v)
        }
    }
    if charCount - digitCount > 1 || charCount - digitCount < -1 { return "" }
    if len(chars) == len(digits) {
        for i, v := range chars {
            res = append(res, []byte{v, digits[i]}...)
        }
    } else if len(chars) > len(digits) {
        for i, v := range chars {
            res = append(res, v)
            if i < len(digits) {
                res = append(res, digits[i])
            }
        }
    } else {
        for i, v := range digits {
            res = append(res, v)
            if i < len(chars) {
                res = append(res, chars[i])
            }
        }
    }
    return string(res)
}