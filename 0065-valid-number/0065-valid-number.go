func isNumber(s string) bool {
    // optionalSign _ 1. / 1.1 / .1 / 1 (number) _ e/E _ optional -/+ _ at least 1 decimal
    validNumber := "^[+-]?((([0-9]+\\.[0-9]*)|(\\.[0-9]+))|([0-9]+))([eE][-+]?[0-9]+)?$"
    // 1. 0 or 1 -/1
    // 2. at least 1 digit decimal
    // 3. at least 1 digit + dot + any digit | dot + at least 1 digit | digit
    // 4. matches e or E
    // 5. -/+ at least 1
    // 6. at least 1 decimal
    // 7. 4, 5 & 6 appear 0 or 1 time
    // 8. end
    reg, err := regexp.Compile(validNumber)
    if err != nil {
        fmt.Println(err)
    }
    
    return reg.MatchString(s)
}