type Heap []int

func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i] > h[j] }
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *Heap) Push(v interface{}) {
	*h = append(*h, v.(int))
}

func (h *Heap) Pop() interface{} {
	cur := *h
	e := cur[len(cur)-1]
	*h = cur[:len(cur)-1]
	return e
}

func kthSmallest(matrix [][]int, k int) int {
	maxHeap := &Heap{}
	heap.Init(maxHeap)

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			heap.Push(maxHeap, matrix[i][j])
			if len(*maxHeap) > k {
				heap.Pop(maxHeap)
			}
		}
	}

	if len(*maxHeap) < k {
		return -1
	}

	return heap.Pop(maxHeap).(int)
}