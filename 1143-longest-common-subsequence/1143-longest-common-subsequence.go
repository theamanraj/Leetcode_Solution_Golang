func max(a, b int) int {
if a > b {
return a
}
return b
}
func longestCommonSubsequence(text1 string, text2 string) int {
var v [][]int
for i := 0; i < len(text1) + 1; i++{
v = append(v, make([]int, len(text2) + 1))
}

for i := 1; i <= len(text1); i++ {
    for j := 1; j <= len(text2); j++ {
        v[i][j] = max(v[i - 1][j], v[i][j - 1])
        if text1[i - 1] == text2[j - 1] {
            v[i][j] = max(v[i][j], v[i - 1][j - 1] + 1)
        } 
    }
}
return v[len(text1)][len(text2)]
}