func fullJustify(words []string, maxWidth int) []string {
    result := []string{}
    tokens := []string{}
    charCount := 0
    
    for len(words) > 0 {
        if charCount+len(tokens)+len(words[0]) <= maxWidth {
            tokens = append(tokens, words[0])
            charCount += len(words[0])
            
            words = words[1:]
            continue
        }
        
        b := strings.Builder{}
        b.Grow(maxWidth)
        
        b.WriteString(tokens[0])
        tokens = tokens[1:]

        totalSpace := maxWidth-charCount

        if len(tokens) == 0 {
            b.WriteString(strings.Repeat(" ", totalSpace))
        }

        for len(tokens) > 0 {
            space := totalSpace/len(tokens)

            if totalSpace%len(tokens) != 0 {
                space++
            }
            
            totalSpace -= space

            b.WriteString(strings.Repeat(" ", space))
            b.WriteString(tokens[0])
            tokens = tokens[1:]
        }
        
        result = append(result, b.String())
        charCount = 0
    }
    
    if len(tokens) > 0 {
        b := strings.Builder{}
        
        b.WriteString(strings.Join(tokens," "))
        b.WriteString(strings.Repeat(" ", maxWidth-b.Len()))
        
        result = append(result, b.String())
    }
    
    return result
}