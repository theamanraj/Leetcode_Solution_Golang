func judgePoint24(nums []int) bool {
	size := len(nums)
	if size == 0 {
		return false
	}

	list := make([]float64, size)
	for i := 0; i < size; i++ {
		list[i] = float64(nums[i])
	}

	return helper(list)
}

func helper(list []float64) bool {
	if len(list) == 1 {
		if math.Abs(list[0]-24.0) < 0.001 {
			return true
		}
		return false
	}

	for i := 0; i < len(list); i++ {
		for j := i + 1; j < len(list); j++ {

			for _, c := range compute(list[i], list[j]) {

				next := make([]float64, 0)
				next = append(next, c)

				for k := 0; k < len(list); k++ {
					if k == i || k == j {
						continue
					}
					next = append(next, list[k])
				}
				if helper(next) {
					return true
				}
			}
		}
	}

	return false
}

func compute(a float64, b float64) []float64 {
	return []float64{a + b, a - b, b - a, a * b, a / b, b / a}
}
