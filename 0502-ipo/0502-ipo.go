func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	pc := make([][]int, 0, len(profits))
	for i := range profits {
		pc = append(pc, []int{profits[i], capital[i]})
	}
	sort.Slice(pc, func(i, j int) bool {
		return pc[i][1] < pc[j][1]
	})
	hp := &Heap{}
	for index := 0; k > 0 && index < len(pc); {
		if w >= pc[index][1] {
			heap.Push(hp, pc[index])
			index++
			continue
		}
		if hp.Len() == 0 {
			break
		}
		job := heap.Pop(hp).([]int)
		w += job[0]
		k--
	}
	for k > 0 && hp.Len() > 0 {
		job := heap.Pop(hp).([]int)
		w += job[0]
		k--
	}
	return w
}

type Heap [][]int

func (p Heap) Len() int            { return len(p) }
func (p Heap) Less(i, j int) bool  { return p[i][0] > p[j][0] }
func (p Heap) Swap(i, j int)       { p[i], p[j] = p[j], p[i] }
func (p *Heap) Push(i interface{}) { *p = append(*p, i.([]int)) }
func (p *Heap) Pop() interface{}   { v := (*p)[len(*p)-1]; *p = (*p)[:len(*p)-1]; return v }