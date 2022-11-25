var ans []string
func generateParenthesis(n int) []string {
    ans = []string{}
    Make(n, 0, "")
    return ans
}

func Make(l, r int, cur string) {
    if l == 0 && r == 0 {
        ans = append(ans, cur)
        return
    }
    if l == 0 {
        for i := 0; i < r; i++ {
            cur += ")"
        }
        ans = append(ans, cur)
        return
    }
    Make(l-1, r+1, cur+"(")
    if r != 0 {
        Make(l, r-1, cur+")")
    }
}