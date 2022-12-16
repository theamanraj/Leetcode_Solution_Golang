func reverseWords(s string) string {
    var words []string = removeExtraSpaces(strings.Split(s, " "))

    var start, end = 0, len(words) - 1

    for start < end {
        var temp string = words[start]
        words[start] = words[end]
        words[end] = temp

        start += 1
        end -= 1
    }

    return joinStrings(words)
}

func joinStrings(words []string) string {
    return strings.Join(words, " ")
}

func removeExtraSpaces(words []string) []string {
    res := make([]string, 0)

    for _, val := range words {
        if strings.TrimSpace(val) != "" {
            res = append(res, val)
        }
    }


    return res
}