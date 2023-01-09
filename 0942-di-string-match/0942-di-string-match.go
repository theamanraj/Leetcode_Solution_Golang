func diStringMatch(S string) []int {
    N := len(S)
	min := 0
    max := N

    A := make([]int, N + 1)
	for i, c := range S {
		if c == 'I' {
			A[i] = min
			min ++
		} else {
			A[i] = max
			max --
		}
	}

	A[N] = min

	return A
}