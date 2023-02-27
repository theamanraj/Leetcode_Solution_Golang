func removePalindromeSub(s string) int {
    for left, right := 0, len(s)-1; left < right; left, right = left+1, right-1 {
        if s[left] != s[right] {
            return 2
        }
    }
    
    return 1
}