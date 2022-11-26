func makeGood(s string) string {
    l := len(s)
    if l == 1 {
        return s
    }
    
    var res strings.Builder
    
    offset := byte('a' - 'A')
    rec := false
    i:=0
    for i<l {
        if i < l - 1 && (s[i] - s[i+1] == offset || s[i+1] - s[i] == offset) {
            i+=2
            rec = true
        } else {
            res.WriteByte(s[i])
            i++
        }
    }
    
    if rec {
        return makeGood(res.String())
    }

    return res.String()
}