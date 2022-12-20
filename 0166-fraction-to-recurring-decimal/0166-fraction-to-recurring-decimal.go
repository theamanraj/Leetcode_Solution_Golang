func fractionToDecimal(n int, d int) string {
    if n % d == 0 {
        return strconv.Itoa(n / d)
    }

    res := ""
    if float64(n) / float64(d) < 0 {
        res += "-"
    }
    
    n, d = int(math.Abs(float64(n))), int(math.Abs(float64(d)))
    res += strconv.Itoa(n / d) + "."
    n %= d
    
    mark, m := len(res), make(map[int]int)
    for n != 0 {
        if _, ok := m[n]; ok {
            mark = m[n]
            return res[:mark] + "(" + res[mark:] + ")"
        } else {
            m[n] = mark
        }
        
        n *= 10
        res += strconv.Itoa(n / d)
        n %= d
        mark++
    }
    return res
}