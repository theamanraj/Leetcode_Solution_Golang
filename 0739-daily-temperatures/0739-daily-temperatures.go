func dailyTemperatures(T []int) []int {
	stack := []int{} // keep decreasing order
	res := make([]int, len(T))
	for i := len(T) - 1; i >= 0; i-- {
		for len(stack) > 0 && T[stack[len(stack)-1]] <= T[i] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) != 0 {
			res[i] = stack[len(stack)-1] - i
		}

		stack = append(stack, i)
	}
	return res
}