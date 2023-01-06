func calPoints(ops []string) int {
	sum, val, stack := 0, 0, make([]int, 0)
	for _, op := range ops {
		switch op {
		case "C":
			val := stack[len(stack)-1]
			sum -= val
			stack = stack[:len(stack)-1]
		case "D":
			val := stack[len(stack)-1] * 2
			sum += val
			stack = append(stack, val)
		case "+":
			val = stack[len(stack)-1] + stack[len(stack)-2]
			sum += val
			stack = append(stack, val)
		default:
			val, _ := strconv.Atoi(op) // ignore error
			stack = append(stack, val)
			sum += val
		}
	}
	return sum
}