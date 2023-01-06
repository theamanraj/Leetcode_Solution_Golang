func shortestCommonSupersequence(str1 string, str2 string) string {
    rs1 := []rune(str1)
    rs2 := []rune(str2)
    lcs := findLongestCommonSequence(rs1, rs2)
    p1, p2, plcs := 0, 0, 0
    var ret []rune
    for ; plcs < len(lcs); p1, p2, plcs = p1+1, p2+1, plcs+1 {
        for ; rs1[p1] != lcs[plcs]; p1++ {
            ret = append(ret, rs1[p1])
        }
        for ; rs2[p2] != lcs[plcs]; p2++ {
            ret = append(ret, rs2[p2])
        }
        ret = append(ret, rs2[p2])
    }
    ret = append(ret, rs1[p1:]...)
    ret = append(ret, rs2[p2:]...)
    return string(ret)
}

func findLongestCommonSequence(rs1, rs2 []rune) []rune {
    m := len(rs1)
    n := len(rs2)
    dp := make([][][]rune, m + 1)
    for i := 0; i <= m; i++ {
        dp[i] = make([][]rune, n + 1)
        for j := 0; j <= n; j++ {
            if i == 0 || j == 0 {
                dp[i][j] = []rune{}
                continue
            }
            if rs1[i-1] == rs2[j-1] {
                clone := make([]rune, len(dp[i-1][j-1]))
                copy(clone, dp[i-1][j-1])
                dp[i][j] = append(clone, rs1[i-1])
                continue
            }
            if len(dp[i-1][j]) > len(dp[i][j-1]) {
                dp[i][j] = dp[i-1][j]
            } else {
                dp[i][j] = dp[i][j-1]
            }
        }
    }
    return dp[m][n]
}