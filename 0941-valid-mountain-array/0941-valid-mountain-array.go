func validMountainArray(A []int) bool {
	if len(A) < 3 {
		return false
	}
	beforeTop := true
	beforeTopCount := 0
	for i := 1; i < len(A); i++ {
		if A[i] == A[i-1] {
			return false
		}
		if beforeTop {
			if A[i] < A[i-1] {
				beforeTop = false
			} else {
				beforeTopCount++
			}
		} else {
			if A[i] > A[i-1] {
				return false
			}
		}
	}
	return !beforeTop && beforeTopCount > 0
}