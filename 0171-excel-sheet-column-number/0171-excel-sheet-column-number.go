func titleToNumber(columnTitle string) int {
  var ans int

	for _, char := range columnTitle {
		ans *= 26
		ans += int((char - 'A') + 1)
	}

	return ans
}