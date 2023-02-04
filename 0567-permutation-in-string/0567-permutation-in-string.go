func checkInclusion(s1 string, s2 string) bool {
    count, cur := [26]int{}, [26]int{}

    for _, r := range s1 {
        count[int(r-'a')]++
    }
    start := 0
    for i, r := range s2 {
        cur[int(r-'a')]++
        for start <= i && cur[int(s2[start]-'a')] > count[int(s2[start]-'a')] {
            cur[int(s2[start]-'a')]--
            start++
        }
    
        if cur == count {
            return true
        }
    }
    
    return cur == count
}