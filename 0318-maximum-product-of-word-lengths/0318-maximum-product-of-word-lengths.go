func maxProduct(words []string) int {
    letterHashes := make([]int, len(words))
 
    for i, word := range words {
        for _, c := range word {
            letterHashes[i] |= 1 << (c-'a')
        }
    }
    
    result := 0
    
    for i := range words {
        for j := i+1; j < len(words); j++ {
            if letterHashes[i] & letterHashes[j] > 0 {
                continue
            }
            
            result = max(result, len(words[i])*len(words[j]))
        }
    }
    
    return result
}

func max(i, j int) int {
    if i > j {
        return i
    }
    
    return j
}