func largestTimeFromDigits(A []int) string {
    ans := -1
    for i := 0; i < 4; i++ {
        for j := 0 ; j < 4; j ++ {
            if j != i {
                for k := 0; k < 4; k ++ {
                    if k != i && k != j {
                        l := 6 - i - j - k
                        
                        hrs := 10 * A[i] + A[j]
                        mns := 10 * A[k] + A[l]
                        
                        if hrs < 24 && mns < 60 {
                            ans = Max(ans, hrs * 60 + mns)
                        } 
                    }
                }
            }
        }
    }
    if ans >= 0 {
        return fmt.Sprintf("%02d:%02d", ans/60, ans%60)
    } 
    return ""
}

func Max(x, y int) int {
    if x < y {
        return y
    }
    return x
} 