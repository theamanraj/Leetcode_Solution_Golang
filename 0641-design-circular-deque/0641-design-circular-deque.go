type MyCircularDeque struct {
    q []int // the Deque
    k int // the lenght of the Deque
}

func Constructor(k int) MyCircularDeque {
    return MyCircularDeque{[]int{}, k}
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	/*
	If the lenght of the deque is not equal to k we can add to the front of the array and return true
	*/

    if len(this.q) == this.k { return false } 
    
    this.q = append(this.q[:0], append([]int{value}, this.q[0:]...)...)
    
    return true
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	
	/*
	If the lenght of the deque is not equal to k we can add to the back of the array and return true
	*/

    if len(this.q) == this.k { return false } 
    
    this.q = append(this.q, value)
    
    return true
}

func (this *MyCircularDeque) DeleteFront() bool {

	/*
	If the lenght of the deque is not equal to 0 we can remove from the front of the array and return true
	*/
	
    if len(this.q) == 0 { return false }
    
    this.q = this.q[1:]
    
    return true
}

func (this *MyCircularDeque) DeleteLast() bool {

	/*
	If the lenght of the deque is not equal to 0 we can remove from the end of the array and return true
	*/

    if len(this.q) == 0 { return false }
    
    this.q = this.q[:len(this.q) - 1]
    
    return true
}

func (this *MyCircularDeque) GetFront() int {

	/*
	If the lenght of the deque is not equal to 0 we can return the first element else return -1
	*/

    if len(this.q) == 0 { return -1 }
    
    return this.q[0]
}

func (this *MyCircularDeque) GetRear() int {

	/*
	If the lenght of the deque is not equal to 0 we can return the last element else return -1
	*/

    if len(this.q) == 0 { return -1 }
    
    return this.q[len(this.q) - 1]
}

func (this *MyCircularDeque) IsEmpty() bool {
    return len(this.q) == 0
}

func (this *MyCircularDeque) IsFull() bool {
    return len(this.q) == this.k
}