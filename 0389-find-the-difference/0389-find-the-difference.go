func findTheDifference(s string, t string) byte {
    
    myMap := make(map[byte]int, len(s))
    
    for i:=0; i<len(s) ; i++ {
        if val, ok := myMap[s[i]]; ok {
            myMap[s[i]] = val +1
        }else {
             myMap[s[i]] = 1
        }
    }
    
    
    for i:=0; i<len(t); i++ {
        if val, ok := myMap[t[i]]; ok {
            if val == 0 {
                return byte(t[i])
            }
            myMap[t[i]] = val-1
        }else {
            return byte(t[i])
        }
    }
    return 0
    
}