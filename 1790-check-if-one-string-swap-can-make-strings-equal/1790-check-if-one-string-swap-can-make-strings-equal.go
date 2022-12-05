func areAlmostEqual(s1 string, s2 string) bool {
    diff := 0
    freq1 := make([]int, 26)
    freq2 := make([]int, 26)
    
    for i := 0; i < len(s1); i++ {
        if s1[i] != s2[i] {
            diff++
        }
        if diff > 2 {
            return false
        }
        freq1[s1[i]-'a']++
        freq2[s2[i]-'a']++
    }
    
    for i := 0; i < 26; i++ {
        if freq1[i] != freq2[i] {
            return false
        }
    }
    
    return true
}