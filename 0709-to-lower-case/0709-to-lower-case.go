func toLowerCase(s string) string {
    x := []byte(s)
    for i:=0;i<len(x);i++ {
        if 'A' <= x[i] && x[i] <= 'Z' {
            x[i] += 'a' - 'A'
        }  
    }
    return string(x)
}