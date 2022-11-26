func isIsomorphic(s string, t string) bool {
    
    if len(s) != len(t) {
        return false
    }
    map1 := make(map[byte]byte)
    map2 := make(map[byte]bool)
    for i:=0; i<len(s);i++ {
        
        value, ok := map1[s[i]]
        if ok {
            if value != t[i] {
                return false
            }
        } else {
            if _, ok := map2[t[i]]; ok {
                return false
            }
            map1[s[i]] = t[i]
            map2[t[i]] = true
        }
    }
    return true
}