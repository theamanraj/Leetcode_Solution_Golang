func isMonotonic(A []int) bool {

	countA := 0
	countB := 0

	if len(A) == 1 {
		return true
	}
	for i := 1; i < len(A); i++ {
		if A[i-1] <= A[i] {
			countA++
		}
		if A[i-1] >= A[i] {
			countB++
		}
		if countA == len(A)-1 || countB == len(A)-1 {
			return true
		}
	}

	return false

}