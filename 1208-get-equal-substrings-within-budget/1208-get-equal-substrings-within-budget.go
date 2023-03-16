func equalSubstring(s string, t string, maxCost int) int {
	var tmp []int
	for i := 0; i < len(s); i++ {
		if s[i] > t[i] {
			tmp = append(tmp, int(s[i]-t[i]))
		} else {
			tmp = append(tmp, int(t[i]-s[i]))
		}
	}
	var i, j, sum, l int = 0, 0, 0, 0
	for j < len(s) {
		sum += tmp[j]
		for sum > maxCost {
			if j-i+1 > l {
				l = j - i + 1
			}
			sum -= tmp[i]
			i++
		}
		j++
	}
	l--
	if j-i > l {
		l = j - i
	}
	if l == -1 {
		return len(s)
	}
	return l
}