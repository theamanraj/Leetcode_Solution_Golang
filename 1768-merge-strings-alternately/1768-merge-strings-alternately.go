func isGreater (w1, w2 string) bool {
    if len(w1) > len(w2) {
        return true
    }
    return false
}

func sliceString (s string, index int) string {
    res := ""
    for i:=index;i<len(s);i++{
        res = res + string(s[i])
    }
    return res
}

func mergeAlternately(word1 string, word2 string) string {
    word := ""
    i := 0
    j := 0 
    for {
        if i<len(word1) && j<len(word2) {
            word = word + string(word1[i]) + string(word2[j])
            i++
            j++
            continue
        }
        break
    }
    
    if isGreater(word1, word2) {
        return word + sliceString(word1, i)
    } 
    
    return word + sliceString(word2, j)
}