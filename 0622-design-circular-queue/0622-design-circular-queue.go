type MyCircularQueue struct {
    values []int
    noOfElementsPresent int
    size int
    front int
    rear int
}


func Constructor(k int) MyCircularQueue {
    return  MyCircularQueue{
        values :make([]int,k),
        noOfElementsPresent : 0,
        size : k,
        front:0,
        rear:-1,
    } 
}


func (this *MyCircularQueue) EnQueue(value int) bool {
    if this.noOfElementsPresent == this.size{
        return false
    }
    this.rear++;
    if this.rear >= this.size{
        this.rear = this.rear%this.size
    }
    this.values[this.rear] = value
    this.noOfElementsPresent++
    return true
}


func (this *MyCircularQueue) DeQueue() bool {
    if this.noOfElementsPresent == 0{
        return false
    }
    this.front++
    this.front = this.front%this.size
    this.noOfElementsPresent--
    return true
}


func (this *MyCircularQueue) Front() int {
    if this.noOfElementsPresent == 0{
        return -1
    }
    return this.values[this.front];
}


func (this *MyCircularQueue) Rear() int {
    if this.noOfElementsPresent == 0{
        return -1
    }
    return this.values[this.rear];
}


func (this *MyCircularQueue) IsEmpty() bool {
    if this.noOfElementsPresent == 0{
        return true
    }
    return false
}


func (this *MyCircularQueue) IsFull() bool {
    if this.noOfElementsPresent == this.size{
        return true
    }
    return false
}