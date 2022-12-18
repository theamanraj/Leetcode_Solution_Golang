func isLetter(b byte) bool {
    return (b >= 'A' &&  b <= 'Z') || (b >= 'a' && b <= 'z') 
}

func reverseOnlyLetters(s string) string {
    x := []byte(s)
    i, j := 0, len(x)-1
    for i < j {
        if !isLetter(x[i]) {
            i++
        }
        
        if !isLetter(x[j]) {
            j--
        }
        
        if isLetter(x[i]) && isLetter(x[j]){
            x[i], x[j] = x[j], x[i]
            i++
            j--
        }  
    }
    return string(x)
}