func swimInWater(grid [][]int) int {
	step := make([][]bool, len(grid))
	for i := range step {
		step[i] = make([]bool, len(grid[0]))
	}
	currentHeight := grid[0][0]
	step[0][0] = true
	h := &Heap{}
	heap.Push(h, []int{0, 0, grid[0][0]})
	for {
		elem := heap.Pop(h).([]int)
		// refresh current height
		if elem[2] > currentHeight {
			currentHeight = elem[2]
		}
		// reached grid[n-1][n-1]
		if elem[0] == len(grid)-1 && elem[1] == len(grid)-1 {
			return currentHeight
		}
		// try up
		if r, c := elem[0] - 1, elem[1]; r >= 0 && step[r][c] == false {
			heap.Push(h, []int{r, c, grid[r][c]})
			step[r][c] = true
		}
		// try down
		if r, c := elem[0] + 1, elem[1]; r < len(grid) && step[r][c] == false {
			heap.Push(h, []int{r, c, grid[r][c]})
			step[r][c] = true
		}
		// try left
		if r, c := elem[0], elem[1] - 1; c >= 0 && step[r][c] == false {
			heap.Push(h, []int{r, c, grid[r][c]})
			step[r][c] = true
		}
		// try right
		if r, c := elem[0], elem[1] + 1; c < len(grid) && step[r][c] == false {
			heap.Push(h, []int{r, c, grid[r][c]})
			step[r][c] = true
		}
	}
}

type Heap [][]int

func (p Heap) Len() int            { return len(p) }
func (p Heap) Less(i, j int) bool  { return p[i][2] < p[j][2] }
func (p Heap) Swap(i, j int)       { p[i], p[j] = p[j], p[i] }
func (p *Heap) Push(i interface{}) { *p = append(*p, i.([]int)) }
func (p *Heap) Pop() interface{}   { v := (*p)[len(*p)-1]; *p = (*p)[:len(*p)-1]; return v }