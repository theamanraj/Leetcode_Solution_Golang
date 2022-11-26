func reverseVowels(s string) string {
    bytes := []byte(s)
    i := 0
    vowels := map[byte]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true, 'A': true, 'E': true, 'O': true, 'I': true, 'U': true, }
    j := len(bytes) - 1
    for i < j{
        for i < j{
            if _, ok := vowels[bytes[i]]; !ok{
                i += 1
            } else{
                break
            }
        }
        for i < j{
            if _, ok := vowels[bytes[j]]; !ok{
                j -= 1
            } else{
                break
            }
        }
        bytes[i], bytes[j] = bytes[j], bytes[i]
        i += 1
        j -= 1
    }
    return string(bytes)
}