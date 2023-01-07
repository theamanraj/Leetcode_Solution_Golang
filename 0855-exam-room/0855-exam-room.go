type ExamRoom struct {
	N         int
	positions []int
}

func Constructor(N int) ExamRoom {
	return ExamRoom{N: N, positions: []int{}}
}

func (this *ExamRoom) Seat() int {
	if len(this.positions) == 0 {
		this.positions = append(this.positions, 0)
		return 0
	}
	p, mx := -1, -1
	if this.positions[0] > 0 {
		p, mx = 0, this.positions[0]-0
	}
	for i := 0; i < len(this.positions)-1; i++ {
		if (this.positions[i+1]-this.positions[i])/2 > mx {
			p, mx = this.positions[i]+(this.positions[i+1]-this.positions[i])/2, (this.positions[i+1]-this.positions[i])/2
		}
	}
	if this.positions[len(this.positions)-1] < this.N-1 {
		if this.N-1-this.positions[len(this.positions)-1] > mx {
			p, mx = this.N-1, this.N-1-this.positions[len(this.positions)-1]
		}
	}

	for i := 0; i < len(this.positions); i++ {
		if this.positions[i] > p {
			l := len(this.positions)
			this.positions = append(this.positions, 0)
			copy(this.positions[i+1:], this.positions[i:l])
			this.positions[i] = p
			return p
		}
	}
	this.positions = append(this.positions, p)
	return p
}

func (this *ExamRoom) Leave(p int) {
	for i := 0; i < len(this.positions); i++ {
		if this.positions[i] == p {
			this.positions = append(this.positions[:i], this.positions[i+1:]...)
			return
		}
	}
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}