func movesToStamp(stamp string, target string) []int {
    temp, ans := []byte(target), []int{}
    m, n, c := len(stamp), len(target), 0
    for c < n {
        stamped := c
        for i := 0; i < n - m + 1; i++ {
            cur := helper(temp[i:i + m], stamp)
            if cur > 0 {
                c += cur
                for j := i; j < i + m; j++ {
                    temp[j] = '?'
                }
                ans = append(ans, i)
            }
        }
        if stamped == c { return []int{} }
    }
    for l, r := 0, len(ans) - 1; l < r; l, r = l + 1, r - 1 { ans[l], ans[r] = ans[r], ans[l] }
    return ans
}

func helper(tmp []byte, stamp string) int {
    s := 0
    for i := 0; i < len(tmp); i++ {
        if tmp[i] == '?' {
            continue
        } else if tmp[i] == stamp[i] {
            s++
        } else {
            return 0
        }
    }
    return s
}