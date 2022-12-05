func interpret(command string) string {
    output := []byte{}
    
    for i := 0; i < len(command); i++ {
        ch := command[i]
        if ch == 'G' {
            output = append(output, ch)
            continue
        }

        if ch == '(' {
            i++
            if command[i] == ')' {
                output = append(output, 'o')
            } else {
                output = append(output, 'a')
                i++
                output = append(output, 'l')
                i++
            }
        }
    }
    
    return string(output)
}