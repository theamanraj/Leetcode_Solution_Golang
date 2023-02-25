func isIdealPermutation(A []int) bool {
    for i := 0; i < len(A); i++ {
        if abs(A[i] - i) > 1 {
            return false
        }
    }
    return true
}

func abs(a int) int {
    if a < 0 {
        return -a 
    } else {
        return a
    }
}