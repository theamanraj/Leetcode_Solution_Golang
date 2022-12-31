type Solution struct {
	MaxLength     int
	ValRow        int
	canNotRandMap []int
}

func Constructor(n_rows int, n_cols int) Solution {
	n_rows, n_cols = n_cols, n_rows
	return Solution{MaxLength: n_rows*n_cols - 1, canNotRandMap: []int{}, ValRow: n_rows}
}

func (this *Solution) Flip() []int {
	row, col, flag, randInt := 0, 0, false, 0
	for true {
		randInt = rand.Intn(this.MaxLength + 1)
		flag = false
		for _, v := range this.canNotRandMap {
			if v == randInt {
				flag = true
				break
			}
		}
		if flag {
			continue
		}
		this.canNotRandMap = append(this.canNotRandMap, randInt)
		break
	}
	col = randInt % this.ValRow
	row = (randInt - col) / this.ValRow
	return []int{row, col}

}

func (this *Solution) Reset() {
	this.canNotRandMap = []int{}
}



/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(m, n);
 * param_1 := obj.Flip();
 * obj.Reset();
 */