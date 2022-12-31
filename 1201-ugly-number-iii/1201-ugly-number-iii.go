func nthUglyNumber(n int, a int, b int, c int) int {
	l, r := int64(1), int64(2000000000)
	lcmAB := lcm(a, b)
	lcmAC := lcm(a, c)
	lcmBC := lcm(b, c)
	lcmABC := lcm(lcmAB, c)
	for l <= r {
		value := (l+r)>>1
		va, vb, vc := value/int64(a), value/int64(b), value/int64(c)
		vab, vbc, vac := value/int64(lcmAB), value/int64(lcmBC), value/int64(lcmAC)
		vabc := value/int64(lcmABC)
		// total: va + vb + vc
		// lcm of AB/BC/AC(except lcm of ABC) count twice
		// lcm of ABC count thrice
		curr := va+vb+vc-(vab-vabc)-(vbc-vabc)-(vac-vabc)-2*vabc
		if curr == int64(n) {
			for {
				if value % int64(a) == 0 || value % int64(b) == 0 || value % int64(c) == 0 {
					return int(value)
				}
				value--
			}
		}
		if curr > int64(n) {
			r = value-1
		} else {
			l = value+1
		}
	}
	return 0
}

// Least Common Multiple
func lcm(v1, v2 int) int {
	mul := v1 * v2
	if v1 < v2 {
		v1, v2 = v2, v1
	}
	for v1 % v2 != 0 {
		v1, v2 = v2, v1 % v2
	}
	return mul / v2
}