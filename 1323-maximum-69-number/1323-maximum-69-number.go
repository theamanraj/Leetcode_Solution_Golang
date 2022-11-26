func maximum69Number (num int) int {
    runes := make([]rune, 0)
    s := strconv.Itoa(num)
    changeOnce := false
    for _, r := range s {
        if r == '6' && !changeOnce {
            runes = append(runes, '9')
            changeOnce = true
        } else {
            runes = append(runes, r)
        }
    }
    
    newNum, _ := strconv.Atoi(string(runes))
    return newNum
}
