func minStickers(stickers []string, target string) int {
    n := len(target)
    dp := make([]int,1<<uint(n))
    for i:=1;i<len(dp);i++ {
        dp[i] = -1
    }
    
    for state :=0;state < len(dp);state++ {
        if dp[state] == -1 {
            continue
        }
        for _,sticker := range stickers {
            now := state
            for _,letter := range sticker {
                for i:=0;i<n;i++ {
                    if (now >> uint(i) & 1) == 1 {
                        continue
                    }
                    
                    if target[i] == byte(letter) {
                        now = (now | 1 << uint(i))
                        break
                    }
                }
            }
            
            if dp[now] == -1 || dp[now] > dp[state] +1 {
                dp[now] = dp[state]+1
            }
        }
        
    }
    return dp[1<<uint(n)-1]
}