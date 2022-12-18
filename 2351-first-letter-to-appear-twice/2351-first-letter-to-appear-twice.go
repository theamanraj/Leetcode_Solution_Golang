func repeatedCharacter(s string) byte {
    sMap := make(map[byte]int)
    
    for i := 0; i < len(s); i++ {
        if count, ok := sMap[s[i]]; ok && count >= 1{
            return s[i]
        }
        
        sMap[s[i]] += 1
    }
    
    return byte(' ')
}