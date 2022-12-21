func wordBreak(s string, wordDict []string) bool {
    dp := make([]int, len(s))
    dp[0] = 1
    if len(s) == 0 {
        return false
    }
    for i := 0; i < len(s); i++ {
        if dp[i] == 0 {
            continue
        }
        for _, v := range wordDict {
            if len(s) >= i + len(v) {
                if s[i:i+len(v)] == v {
                    if i + len(v) == len(s) {
                        return true
                    }
                    dp[i + len(v)] = 1
                    // fmt.Println(i, i + len(v), len(s))  
                }
            }
        }
    }
    return false
}