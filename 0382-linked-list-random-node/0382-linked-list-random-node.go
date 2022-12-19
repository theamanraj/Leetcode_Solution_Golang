type Solution struct {
    arr []int
}

/** @param head The linked list's head.
        Note that the head is guaranteed to be not null, so it contains at least one node. */
func Constructor(head *ListNode) Solution {
    arr := []int{}
    for head != nil {
        arr = append(arr, head.Val)
        head = head.Next
    }
    return Solution{ arr }
}

/** Returns a random node's value. */
func (this *Solution) GetRandom() int {
    return this.arr[rand.Intn(len(this.arr))]
}