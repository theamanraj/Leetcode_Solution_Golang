func smallestRangeI(A []int, K int) int {

	if len(A) < 2 {
		return 0
	}

	if len(A) == 2 {
		diff := int(math.Abs(float64(A[0] - A[1])))
		if K == 0 {
			return diff
		}
		if diff <= K*2 {
			return 0
		} else {
			return diff - K*2
		}
	}

	if len(A) > 2 {

		min := A[0]
		max := A[0]

		for _, v := range A {
			if v < min {
				min = v
			}
			if v > max {
				max = v
			}
		}
        
		diff := max - min

		if K == 0 {
			return diff
		}
		if diff <= K*2 {
			return 0
		} else {
			return diff - K*2
		}
	}

	return 0

}