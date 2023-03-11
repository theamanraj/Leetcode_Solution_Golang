func minIncrementForUnique(A []int) int {
	sort.Ints(A)
	numIncs := 0
	l := len(A)
	for i := 1; i < l; i += 1 {
		if A[i] <= A[i-1] {
			numIncs += A[i-1] + 1 - A[i]
			A[i] = A[i-1] + 1
		}
	}
	return numIncs
}