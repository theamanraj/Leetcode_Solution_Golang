func maxArea(height []int) int {
	length := len(height)
	if length == 0 || length == 1 {
		return 0
	}

	left := 0
	right := length - 1

	result := 0
	for left != right {
		result = max(result, (right-left)*min(height[right], height[left]))
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}