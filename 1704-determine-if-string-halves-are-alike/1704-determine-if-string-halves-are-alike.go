func halvesAreAlike(s string) bool {
    a, b, n := 0, 0, len(s)
    vowels := map[byte]bool{
        'a':true,
        'e':true,
        'i':true,
        'o':true,
        'u':true,
        'A':true,
        'E':true,
        'I':true,
        'O':true,
        'U':true,
    }
    for i := 0; i < n / 2; i++ { if vowels[s[i]] { a++ } }
    for i := n / 2; i < n; i++ { if vowels[s[i]] { b++ } }
    return a == b
}