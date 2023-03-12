type minHeap []int

func (h minHeap) Len() int { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h minHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *minHeap) Pop() interface{} {
    min := (*h)[len(*h) - 1]
    *h = (*h)[:len(*h) - 1]
    return min
}

func mergeKLists(lists []*ListNode) *ListNode {
    pq := &minHeap{}
    heap.Init(pq)
    for _, v := range lists {
        for v != nil {
            heap.Push(pq, v.Val)
            v = v.Next
        }
    }
    dummy := new(ListNode)
    head := dummy
    for pq.Len() > 0 {
        head.Next = &ListNode{
            Val: heap.Pop(pq).(int),
            Next: nil,
        }
        head = head.Next
    }
    return dummy.Next
}