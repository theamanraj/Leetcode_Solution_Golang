type Route struct {
	cityIdx  int32
	distance int16
}

func min(a, b int16) int16 {
	if a <= b {
		return a
	}
	return b
}

func minScore(n int, roads [][]int) int {
	cities := make([][]Route, n+1)
	for _, road := range roads {
		cities[road[0]] = append(cities[road[0]], Route{
			cityIdx:  int32(road[1]),
			distance: int16(road[2]),
		})
		cities[road[1]] = append(cities[road[1]], Route{
			cityIdx:  int32(road[0]),
			distance: int16(road[2]),
		})
	}
	result := int16(10000)
	visited := make([]bool, n+1)
	visited[1] = true
	stack := []int32{1}
	for len(stack) != 0 {
		cityIdx := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		for _, route := range cities[cityIdx] {
			if !visited[route.cityIdx] {
				visited[route.cityIdx] = true
				stack = append(stack, route.cityIdx)
			}
			result = min(result, route.distance)
		}
	}
	return int(result)
}