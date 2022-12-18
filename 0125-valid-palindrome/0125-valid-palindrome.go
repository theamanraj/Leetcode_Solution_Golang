func isPalindrome(s string) bool {
    l := 0
    r := len(s) - 1
    for l < r {
        for l < r && !unicode.IsLetter(rune(s[l])) && !unicode.IsNumber(rune(s[l])) { l++ }
        for l < r && !unicode.IsLetter(rune(s[r])) && !unicode.IsNumber(rune(s[r])) { r-- }
        if unicode.ToLower(rune(s[l])) != unicode.ToLower(rune(s[r])) { return false }
        l++
        r--
    }
    return true
}