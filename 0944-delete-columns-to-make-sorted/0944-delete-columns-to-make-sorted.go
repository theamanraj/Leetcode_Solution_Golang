func minDeletionSize(A []string) int {

	length := len(A[0])
	ss := make([][]rune, length)
    count := 0

	for i, v := range ss {
		for _, V := range A {

			v = append(v, rune(V[i]))
		}
		ss[i] = v
	}
	for _, v := range ss {

		for I := 0; I < len(v)-1; I++ {

			if v[I] > v[I+1] {
				count++
				break
			}
		}
	}
	return count
}