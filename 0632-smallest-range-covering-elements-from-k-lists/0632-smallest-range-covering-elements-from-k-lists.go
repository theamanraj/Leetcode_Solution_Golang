type HeapNode struct {
    Val int 
    From int //from which list
}

func buildMinHeap(heap []HeapNode) {
    if len(heap) < 2 {
        return
    }
    for i := len(heap)/2-1; i >= 0; i-- {
        shitDown(heap, i)
    }
}

func shitDown(heap []HeapNode, idx int) {
    if len(heap) < 2 {
        return
    }
    
    for idx < len(heap) {
        i := 2 * idx + 1 //left child
        if i >= len(heap) { 
            break
        }
        
        min := i
        if j := i+1; j < len(heap) && heap[j].Val < heap[i].Val { //right child exist && right child less than left child
            min = j
        }
        
        if heap[min].Val >= heap[idx].Val {
            break
        }
        heap[idx], heap[min] = heap[min], heap[idx]
        idx = min
    }
}

func shitUp(heap []HeapNode, idx int) {
    if len(heap) < 2 {
        return
    }
    
    for idx > 0 {
        p := (idx - 1) / 2
        if heap[p].Val <= heap[idx].Val { 
            break
        }
        heap[p], heap[idx] = heap[idx], heap[p]
        idx = p
    }
}

func deleteMin(heap []HeapNode) HeapNode {
    min := heap[0]
    heap[0], heap[len(heap)-1] = heap[len(heap)-1], heap[0]
    shitDown(heap[:len(heap)-1], 0)
    return min
}

func smallestRange(nums [][]int) []int {
    if len(nums) == 0 {
        return nil
    }
    
    res := make([]int, 2)
    var min, max, length, curIdx int 
    var top HeapNode //top of min-heap
    max = nums[0][0]
    length = math.MaxInt32
    heap := make([]HeapNode, len(nums)) //min-heap
    indices := make([]int, len(nums)) //current index for each list
    for i := 0; i < len(nums); i++ {
        heap[i] = HeapNode{
            Val: nums[i][0],
            From: i,
        }
        if nums[i][0] > max {
            max = nums[i][0]
        }
    }
    buildMinHeap(heap) //init min-heap
    
    for {
        top = deleteMin(heap)
        heap = heap[:len(heap)-1]
        min = top.Val
        curIdx = top.From
        if max - min < length { //found smaller range
            res[0] = min
            res[1] = max
            length = max - min
        }
        
        if indices[curIdx] == len(nums[curIdx])-1 { 
            break
        }
        
        heap = append(heap, HeapNode{
            Val: nums[curIdx][indices[curIdx]+1],
            From: curIdx,
        })
        indices[curIdx]++
        shitUp(heap, len(heap)-1)
        if nums[curIdx][indices[curIdx]] > max {
            max = nums[curIdx][indices[curIdx]]
        }                                    
    }
    
    return res   
}