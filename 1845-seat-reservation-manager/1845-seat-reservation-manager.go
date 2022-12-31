type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type SeatManager struct {
	seatsSlice IntHeap
}

func Constructor(n int) SeatManager {
	seatsSlice := IntHeap{}

	for i := 1; i <= n; i++ {
		seatsSlice = append(seatsSlice, i)
	}

	return SeatManager{
		seatsSlice: seatsSlice,
	}
}

func (this *SeatManager) Reserve() int {
	res := heap.Pop(&this.seatsSlice)

	return res.(int)
}

func (this *SeatManager) Unreserve(seatNumber int) {
	heap.Push(&this.seatsSlice, seatNumber)
}