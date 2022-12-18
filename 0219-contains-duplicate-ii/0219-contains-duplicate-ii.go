func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int)

	for i1, n := range nums {
		if i2, ok := m[n]; ok {
			if i1-i2 <= k {
				return true
			}
		}
		m[n] = i1
	}

	return false
}