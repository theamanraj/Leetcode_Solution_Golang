type t struct {
	s int
	e int
	p int
}

func jobScheduling(st []int, et []int, p []int) int {
	r := make([]t, len(st))
	for i := range st {
		r[i] = t{s: st[i], e: et[i], p: p[i]}
	}

	sort.Slice(r, func(i, j int) bool { return r[i].s <= r[j].s })

	var dp func(int) int
	dp = func(i int) int {
		if i >= len(r) {
			return 0
		}

		a := r[i].p
		for j := i + 1; j < len(r); j++ {
			if r[i].e <= r[j].s {
				a += dp(j)
				break
			}
		}
		na := dp(i + 1)
		if na > a {
			return na
		}
		return a
	}
	dp = c(dp)
	return dp(0)
}

func c(f func(int) int) func(int) int {
	var m = make(map[int]int)
	d := func(arg int) int {
		if val, ok := m[arg]; ok {
			return val
		}
		val := f(arg)
		m[arg] = val
		return val
	}
	return d
}