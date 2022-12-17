func div(i, j int) int {
	if i%j == 0 {
		return i / j
	} else {
		return i/j + 1
	}
}
func helper(nums []int, divisor int) int {
	var res int = 0
	for _, v := range nums {
		res += div(v, divisor)
	}
	return res
}
func smallestDivisor(nums []int, threshold int) int {
	var l, r int = 1, 1000000
	for l < r {
		var m int = (l + r) / 2
		var tmp int = helper(nums, m)
		if tmp > threshold {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}
