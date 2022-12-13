func divide(dividend int, divisor int) int {
	ret := dividend / divisor
	if ret > 2147483647 {
		return 2147483647
	}
	if ret < -2147483648 {
		return -2147483648
	}
	return ret
}