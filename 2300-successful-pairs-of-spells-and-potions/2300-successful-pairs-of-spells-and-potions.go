func successfulPairs(spells []int, potions []int, success int64) []int {
	sort.Ints(potions)
	n := len(spells)
	m := len(potions)
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		left := 0
		right := m

		for left < right {
			mid := left + (right - left) / 2
			if int64(potions[mid]*spells[i]) >= success {
				right = mid 
			} else {
				left = mid + 1
			}
		}
		ans[i]= m - left
	}
	return ans
}