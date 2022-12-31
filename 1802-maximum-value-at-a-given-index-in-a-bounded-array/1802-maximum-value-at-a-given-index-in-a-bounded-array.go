func check(n, index, val, maxSum int) bool {
	left := index + 1
	right := n - 1 - index
	var leftSum, rightSum int
	if left >= val {
		leftSum = left - val + val*(val+1)/2
	} else {
		start := val - left + 1
		leftSum = val*(val+1)/2 - start*(start+1)/2 + start
	}

	val = val - 1
	if right >= val {
		rightSum = right - val + val*(val+1)/2
	} else {
		end := val - right + 1
		rightSum = val*(val+1)/2 - end*(end+1)/2 + end
	}
	return leftSum+rightSum <= maxSum
}

func maxValue(n int, index int, maxSum int) int {
	r := int(1e9)+1
	l := 1
	for l < r {
		mid := (l + r) / 2
		if check(n, index, mid, maxSum) {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l - 1
}