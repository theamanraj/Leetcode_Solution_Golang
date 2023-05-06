func reorderLogFiles(logs []string) []string {
    var letters, digits, result []string
    
    for _, v := range logs {
        tokens := strings.Split(v, " ")
        
        if unicode.IsNumber(rune(tokens[1][0])) {
            digits = append(digits, v)
        } else {
            letters = append(letters, v)
        }
    }
    
    sort.Slice(letters, func(i, j int) bool {
        first := strings.SplitN(letters[i], " ", 2)
        second := strings.SplitN(letters[j], " ", 2)
        
        if strings.Compare(first[1], second[1]) == 0 {
            return strings.Compare(first[0], second[0]) < 0
        }
        
        return strings.Compare(first[1], second[1]) < 0
    })
    
    result = append(result, letters...)
    result = append(result, digits...)
    
    return result
}