func minCostClimbingStairs(cost []int) int {
	cost = append(cost, 0)
	m := make([]int, len(cost))
	for i := range m {
		m[i] = math.MaxInt
	}
	m[0] = 0
	m[1] = 0
	for i := 2; i < len(cost); i++ {
		if m[i-1]+cost[i-1] < m[i] {
			m[i] = m[i-1] + cost[i-1]
		}
		if m[i-2]+cost[i-2] < m[i] {
			m[i] = m[i-2] + cost[i-2]
		}
	}
	//fmt.Println(m)
	return m[len(m)-1]
}