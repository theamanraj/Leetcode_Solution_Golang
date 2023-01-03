func threeEqualParts(A []int) []int {
    var c, i, s1, s2, s3 int
    for i = range A { if A[i] == 1 { c++ } }
    
    if c%3 != 0 || len(A) < 3 { return []int{-1, -1} }
    if c == 0 { return []int{0, len(A)-1} }
    
    for A[s1] != 1 { s1++ }
    s1++
    for s2, i = s1, c/3; i > 0; s2++ { if A[s2] == 1 { i-- } }
    for s3, i = s2, c/3; i > 0; s3++ { if A[s3] == 1 { i-- } }
    for s3 < len(A) && A[s1] == A[s2] && A[s2] == A[s3] {
        s1++
        s2++
        s3++
    }
    if s3 < len(A) { return []int{-1, -1} }
    return []int{s1-1, s2}
}