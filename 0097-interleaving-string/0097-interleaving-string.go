func isInterleave(s1 string, s2 string, s3 string) bool {
    if len(s3) == 0 {
        return true
    }
    if len(s1) == 0 {
        return s2 == s3
    }
    if len(s2) == 0 {
        return s1 == s3
    }
    if len(s3) != len(s1) + len(s2) {
        return false
    }
    var v [][]bool
    for i := 0; i <= len(s1); i++ {
        v = append(v, make([]bool, len(s2) + 1))
    }
    v[0][0] = true
    for i := 0; i <= len(s1); i++ {
        for j := 0; j <= len(s2); j++ {
            if j > 0 && s3[i + j - 1] == s2[j - 1] {
                v[i][j] = v[i][j - 1] 
            } 
            if i > 0 && s3[i + j - 1] == s1[i - 1] {
                v[i][j] = v[i - 1][j] || v[i][j]
            }
        }
    }
    return v[len(s1)][len(s2)]
}