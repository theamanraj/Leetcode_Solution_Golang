type Heap [][3]int

func (p Heap) Len() int            { return len(p) }
func (p Heap) Less(i, j int) bool  { return p[i][0] < p[j][0] }
func (p Heap) Swap(i, j int)       { p[i], p[j] = p[j], p[i] }
func (p *Heap) Push(i interface{}) { *p = append(*p, i.([3]int)) }
func (p *Heap) Pop() interface{}   { v := (*p)[len(*p)-1]; *p = (*p)[:len(*p)-1]; return v }

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	// valid check
	if len(nums1) == 0 || len(nums2) == 0 || k <= 0 {
		return nil
	}
	minHeap := Heap{}
	for i, v := range nums1 {
		minHeap = append(minHeap, [...]int{v+nums2[0], i, 0})
	}
	heap.Init(&minHeap)
	// fix k
	if k > len(nums1) * len(nums2) {
		k = len(nums1) * len(nums2)
	}
	pairs := make([][]int, 0, k)
	for i := 0; i < k; i++ {
		top := heap.Pop(&minHeap).([3]int)
		pairs = append(pairs, []int{nums1[top[1]], nums2[top[2]]})
		if top[2] < len(nums2) - 1 {
			top[2]++
			top[0] = nums1[top[1]]+nums2[top[2]]
			heap.Push(&minHeap, top)
		}
	}
	return pairs
}