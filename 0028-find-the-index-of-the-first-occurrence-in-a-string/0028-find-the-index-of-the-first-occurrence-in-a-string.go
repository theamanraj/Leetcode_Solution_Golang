func strStr(haystack string, needle string) int {
    hayLen, neeLen := len(haystack), len(needle)    
    
    for i:=0;i<hayLen-neeLen+1;i++{
        if haystack[i:i+neeLen]==needle{return i} 
    }
    return -1
}
