func compress(chars []byte) int {
    
    scan := 0
    write := 0
    l := len(chars)
    
    for scan < l {
        count := 0
        chars[write] = chars[scan]
        for scan < l && chars[write] == chars[scan] {
            count++
            scan++
        }
        if count > 1 {
            tmp := fmt.Sprintf("%d", count)
            for _, c := range []byte(tmp) {
                write++
                chars[write] = c
            }
        }
        write++
    }
    return write    
}