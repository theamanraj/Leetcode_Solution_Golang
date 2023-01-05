func expressiveWords(S string, words []string) int {
   
    count := 0
    
    for _, w := range words {
        if stretchy(w, S) {
            count ++
        }
    }
    
    return count    
}

func stretchy(w, s string) bool {
    
    if len(s) == 0 {
        return len(w) == 0
    }
    
    sind := 0
    wind := 0
    
    for wind < len(w) && sind < len(s) {
        
        if w[wind] != s[sind] {
            return false
        }
        
        wdiff := 1
        sdiff := 1
        
        for ; wind + wdiff < len(w) && w[wind + wdiff] == w[wind]; wdiff ++ {  
        }
        
        for ; sind + sdiff < len(s) && s[sind + sdiff] == s[sind]; sdiff ++ {
            
        }
        
        wind += wdiff
        sind += sdiff
        
        if sdiff == wdiff {
            // same number of chars
            continue
        }
        
        if wdiff > sdiff {
            // fewer chars
            return false
        }
        
        if sdiff < 3 {
            return false
        }   
    }
    
    if sind < len(s) || wind<len(w) {
        return false
    }
    
    return true
}