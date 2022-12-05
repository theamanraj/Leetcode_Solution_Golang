func isAlienSorted(words []string, order string) bool {
    byteOrder := map[byte]int{}
    for i:=1;i<len(order); i++{ byteOrder[order[i]]=i}
    
    for i:=1;i<len(words); i++{
        if !isOrder(words[i-1],words[i], byteOrder){ return false   }
    }    
    return true
}


func isOrder( a, b string, byteOrder map[byte]int )bool{    
    for i:=0;i<len(a) && i<len(b); i++{
        if aP, bP := byteOrder[a[i]], byteOrder[b[i]] ; aP>bP{
            return false
        }else if aP < bP{     return true  }
    }
    return len(a)<=len(b)  // if preFix bytes are the same, compare length.
}