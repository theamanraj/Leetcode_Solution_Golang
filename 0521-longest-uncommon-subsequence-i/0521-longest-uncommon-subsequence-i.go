func findLUSlength(a string, b string) int {
    if a == b {
        return -1
    }   
    
    return len(a) - ((len(a) - len(b)) & ((len(a) -len(b)) >> 31 ))
}