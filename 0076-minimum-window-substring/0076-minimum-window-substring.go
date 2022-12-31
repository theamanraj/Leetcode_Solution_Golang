func minWindow(s string, t string) string {
    left, right, minLength := 0, 0, math.MaxInt32
    minStr := ""
    
    tHashMap := make(map[rune]int)
    sHashMap := make(map[rune]int)
    for i := 0; i < 26; i++ {  // as we know the strings are lower and upper case letters from the English alphabet
        tHashMap[rune(int(rune('a'))+i)] = 0
        tHashMap[rune(int(rune('A'))+i)] = 0
        sHashMap[rune(int(rune('a'))+i)] = 0
        sHashMap[rune(int(rune('A'))+i)] = 0  
    }
    
    for _, c := range t {
        tHashMap[rune(c)]++
    } 
    
    for ;left <= right && right < len(s); {
        for ;!satisfies(sHashMap, tHashMap) && right < len(s); right++{ // increment right until the substring of s satifies t
            sHashMap[rune(s[right])]++
        }
        for ;satisfies(sHashMap, tHashMap) && left < right; left++{ // increment left until the substring of s does not satisfy t
            if minLength > right - left {
                minLength = right - left
                minStr = s[left:right]
            }
            sHashMap[rune(s[left])]--
        }
    }
    return minStr   
}

func satisfies(sHashMap, tHashMap map[rune]int) bool {
    for k, v := range tHashMap {
        count, _ := sHashMap[k]
        if count < v {
            return false
        }  
    }
    return true
}