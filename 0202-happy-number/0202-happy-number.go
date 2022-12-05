func isHappy(n int) bool {
	visited := make(map[int]bool)
	for n != 1 {
		if _, ok := visited[n]; ok {
			return false
		}
		visited[n] = true

		m := 0
		temp := n
		for temp > 0 {
			last := temp % 10
			m = m + last*last
			temp = temp / 10
		}
		n = m
	}

	return true
}