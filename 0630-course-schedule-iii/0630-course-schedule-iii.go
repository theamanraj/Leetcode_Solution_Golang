type IntHeap []int

// max heap
func (h IntHeap) Less(i,j int) bool { return h[i] > h[j] }

func (h IntHeap) Swap(i,j int) { h[i], h[j] = h[j], h[i] }

func (h IntHeap) Len() int { return len(h) }

func (h *IntHeap) Push(val interface{}) {
    *h = append(*h, val.(int))
}

// pop the top element for max heap
func (h *IntHeap) Pop() interface{} {
    old := *h
    l := len(*h)
    val := old[l-1]
    *h = old[0:l-1]
    
    return val
}

func scheduleCourse(courses [][]int) int {
    
    h := &IntHeap{}
    heap.Init(h)
    
    // sort courses by the last day, the courses ending on the same day will be handled by max heap
    sort.Slice(courses, func(i,j int) bool {
        return courses[i][1] < courses[j][1]
    })
    
    daysUsed := 0
    for _, course := range courses {
        
        daysUsed += course[0]
        // adding the duration of the current course
        heap.Push(h, course[0])
        
        // if the course is not doable, then pop the course that took the most days
        if daysUsed > course[1] {
            daysUsed -= heap.Pop(h).(int)
        }
    }
    // the final size of the queue is the total courses one can take at max!
    
    return h.Len()
}