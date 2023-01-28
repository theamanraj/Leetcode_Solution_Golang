func solveEquation(equation string) string {
	sides := strings.Split(equation, "=")

	xL, nL := xAndNum("+" + sides[0])
	xR, nR := xAndNum("+" + sides[1])

	if xR == xL {
		if nL == nR {
			return "Infinite solutions"
		}
		return "No solution"
	}
	return fmt.Sprintf("x=%v", (nL-nR)/(xR-xL))
}

func xAndNum(s string) (int, int) {
	x := 0
	num := 0
	for len(s) > 0 {
		i := 1
		for i < len(s) {
			if s[i] == '+' || s[i] == '-' {
				break
			}
			i++
		}
		value := 0
		if s[i-1] != 'x' {
			value, _ = strconv.Atoi(s[1:i])
			if s[0] == '+' {
				num = num + value
			} else {
				num = num - value
			}
		} else {
			if i == 2 {
				value = 1
			} else {
				value, _ = strconv.Atoi(s[1 : i-1])
			}
			if s[0] == '+' {
				x = x + value
			} else {
				x = x - value
			}
		}
		s = s[i:]
	}
	return x, num
}