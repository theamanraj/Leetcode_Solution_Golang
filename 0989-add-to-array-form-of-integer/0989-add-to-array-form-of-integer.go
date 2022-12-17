func addToArrayForm(A []int, K int) []int {
    res := make([]int, len(A)+1)
    kLen := len(strconv.Itoa(K))
    if kLen > len(A) {
        res = make([]int, kLen+1)
    }
    
    for carry, i, k := 0, len(A)-1, len(res)-1; k >= 0; carry, i, k = carry/10, i-1, k-1 {
        if i >= 0 {
            carry += A[i]
        }
        if K > 0 {
            carry += K % 10
            K /= 10
        }
        res[k] = carry % 10
    }
    
    if res[0] == 0 {
        return res[1:]
    }
    return res
}