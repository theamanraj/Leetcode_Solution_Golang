func isSelfCrossing(distance []int) bool {
	distance = append([]int{0, 0, 0, 0}, distance...)
	n := len(distance)
	i := 4
	for i < n && distance[i] > distance[i-2] {
		i++
	}
	if i == n {
		return false
	}

	if distance[i] >= distance[i-2]-distance[i-4] {
		distance[i-1] -= distance[i-3]
	}
	i = i + 1
	for i < n && distance[i] < distance[i-2] {
		i++
	}
	return i != n

}