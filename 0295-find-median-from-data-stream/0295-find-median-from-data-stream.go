import (
    "container/heap"
)

type MyHeap struct {
    data []int
    less func(arr []int, i, j int) bool
}

func (h MyHeap) Len() int { return len(h.data) }
func (h MyHeap) Swap(i, j int) { h.data[i], h.data[j] = h.data[j], h.data[i] }
func (h MyHeap) Less(i, j int) bool { return h.less(h.data, i, j)  }
//heap specific
func (h *MyHeap) Push(x interface{}) { h.data = append(h.data, x.(int)) }
func (h *MyHeap) Peek() interface{} { return h.data[ 0 ] }
func (h *MyHeap) Pop() interface{} {
    old := h.data
    l := len(old)
    ele := old[l - 1]
    h.data = old[:l-1]
    return ele
}

func newHeap(less func(arr []int, i, j int) bool) *MyHeap {
    return &MyHeap {
        data : make([]int, 0),
        less : less,
    }
}

type MedianFinder struct {
    biggerHalf *MyHeap
    smallerHalf *MyHeap 
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
    ascOrder := func(arr []int, i, j int) bool {
        return arr[i] < arr[j]
    }
    
    dscOrder := func(arr []int, i, j int) bool {
        return arr[i] > arr[j]
    }
    
    return MedianFinder{
        biggerHalf: newHeap( ascOrder ),
        smallerHalf: newHeap( dscOrder ),
    }
}

func (this *MedianFinder) AddNum(num int)  {
    if this.biggerHalf.Len() == 0 || this.smallerHalf.Len() == 0 {
        heap.Push(this.biggerHalf, num)
    } else {
        topOfSmaller := this.smallerHalf.Peek().(int)
        if num <= topOfSmaller {
            heap.Push(this.smallerHalf, num)
        } else {
            heap.Push(this.biggerHalf, num)
        }
    }
    
    if this.smallerHalf.Len() > this.biggerHalf.Len() {
        heap.Push( this.biggerHalf, heap.Pop(this.smallerHalf).(int) )
    } else if this.biggerHalf.Len() - this.smallerHalf.Len() >= 2 {
        heap.Push( this.smallerHalf, heap.Pop(this.biggerHalf).(int) )
    }
}

func (this *MedianFinder) FindMedian() float64 {
    if this.biggerHalf.Len() == 0 {
        return 0
    }
    
    if this.biggerHalf.Len() > this.smallerHalf.Len() {
        return float64( this.biggerHalf.Peek().(int) )
    }
    
    left, right := this.smallerHalf.Peek().(int), this.biggerHalf.Peek().(int)
    
    return ( float64(left) + float64(right) ) / 2
}


/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */