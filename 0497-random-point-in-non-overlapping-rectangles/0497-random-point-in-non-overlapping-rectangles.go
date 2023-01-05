type Solution struct {
	magic []int
	rects [][]int
}

func Constructor(rects [][]int) Solution {
	magic := make([]int, 0, len(rects))
	summary := -1
	for _, v := range rects {
		summary += (v[2]-v[0]+1)*(v[3]-v[1]+1)
		magic = append(magic, summary)
	}
	return Solution{
		magic: magic,
		rects: rects,
	}
}

func (s *Solution) Pick() []int {
	choice1 := rand.Intn(s.magic[len(s.magic)-1]+1)
	l, r := 0, len(s.magic)-1
	for l <= r {
		m := (l+r)>>1
		if s.magic[m] < choice1 {
			l = m+1
		} else {
			r = m-1
		}
	}
	index := l
	choice2 := rand.Intn(100000000)
	x := s.rects[index][0] + choice2 % 10000 % (s.rects[index][2] - s.rects[index][0] + 1)
	y := s.rects[index][1] + choice2 / 10000 % (s.rects[index][3] - s.rects[index][1] + 1)
	return []int{x,y}
}