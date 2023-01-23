func uniqueMorseRepresentations(words []string) int {
    table := []string{".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."}
    
    result := make(map[string]bool)
    for _, w := range words {
        var res string
        for _, r  := range w {
            res += table[unicode.ToLower(r)-'a']
        }
        result[res] = true
    }
    
    return len(result)
}