func kthPalindrome(queries []int, intLength int) []int64 {
    half := (intLength - 1) / 2 + 1
    
    start := 1
    for i := 1; i < half; i++ {
       start *= 10 
    }
    
    numPalin := start * 9
    start--
    doubleLast := (intLength & 1) == 0
    
    ans := make([]int64, len(queries)) 
    for i := range queries {
        if queries[i] > numPalin {
            ans[i] = -1
        } else {
            ans[i] = mirror(start + queries[i], doubleLast)
        }
    }
    
    return ans
}

func mirror(v int, doubleLast bool) int64 {
    ans := 0
    p := 1 // power of 10
    for v > 0 {
        d := v % 10 // digit
        v /= 10
        
        if p == 1 { // first
            if doubleLast {
                ans = d * 11
                p = 100
            } else {
                ans = d
                p = 10
            }
        } else {
            ans = 10 * (d * p + ans) + d
            p *= 100
        }
    } 
    
    return int64(ans)
}