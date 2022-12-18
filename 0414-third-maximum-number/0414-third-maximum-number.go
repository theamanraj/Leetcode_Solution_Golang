type MinHeap []int

func (h MinHeap) Len() int {
    return len(h)
}

func (h MinHeap) Swap(i, j int) {
    h[i], h[j] = h[j], h[i]
}

func (h MinHeap) Less(i, j int) bool {
    return h[i] < h[j]
}

func (h *MinHeap) Push(i interface{}) {
    *h = append(*h, i.(int))
}

func (h *MinHeap) Pop() interface{} {
    x := (*h)[len(*h) - 1]
    *h = (*h)[0: len(*h) - 1]
    return x
}

func thirdMax(nums []int) int {
    exist := make(map[int]bool)
    minHeap := make(MinHeap, 0)
    heap.Init(&minHeap)
    for _, n := range nums {
        if _, ok := exist[n]; ok {
            continue
        }
        heap.Push(&minHeap, n)
        exist[n] = true
        if minHeap.Len() > 3 {
            i := heap.Pop(&minHeap).(int)
            delete(exist, i)
        }
    }
    res := 0
    for minHeap.Len() > 0 {
        res = heap.Pop(&minHeap).(int)
        if minHeap.Len() == 2 { return res }
    }
    return res
}