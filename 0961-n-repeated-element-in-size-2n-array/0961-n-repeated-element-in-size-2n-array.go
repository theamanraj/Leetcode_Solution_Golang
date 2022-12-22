func repeatedNTimes(A []int) int {
    var contained []int
	for i := 0; i < len(A); i++ {
		if !contains(contained, A[i]) {
			contained = append(contained, A[i])
		} else {
			return A[i]
		}
	}
	return 0
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}