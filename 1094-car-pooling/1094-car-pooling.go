type MinHeap [][]int

func (h MinHeap) Len() int { return len(h) }
func (h MinHeap) Less(i, j int) bool {
	if h[i][0] != h[j][0] {
		return h[i][0] < h[j][0]
	}
	return h[i][1] < 0
}
func (h MinHeap) Swap(i, j int)           { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(value interface{}) { *h = append(*h, value.([]int)) }
func (h *MinHeap) Pop() interface{} {
	min := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return min
}

func carPooling(trips [][]int, capacity int) bool {
	h := MinHeap{}
	for i := 0; i < len(trips); i++ {
		heap.Push(&h, []int{trips[i][1], trips[i][0]})
		heap.Push(&h, []int{trips[i][2], -trips[i][0]})
	}

	count := 0
	for h.Len() > 0 {
		c := heap.Pop(&h).([]int)
		count += c[1]
		if count > capacity {
			return false
		}
	}
	return true
}

func abs(x int) int {
	if x > 0 {
		return x
	}

	return -x
}