type Node struct {
    val int
    prev, next *Node
}

type MyLinkedList struct {
    length int
    start, end *Node
}


/** Initialize your data structure here. */
func Constructor() MyLinkedList {
    start := &Node{}
    end := &Node{}
    start.next = end
    end.prev = start
    ll := MyLinkedList{length: 0, start: start, end: end}
    return ll
}


/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
    if this.length <= index {
        return -1
    }
    p := this.start
    for i := 0; i <= index; i++ {
        p = p.next
    }
    return p.val
}


/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int)  {
    curtNext := this.start.next
    newNode := &Node{val, nil, nil}
    this.start.next = newNode
    newNode.prev = this.start
    newNode.next = curtNext
    curtNext.prev = newNode
    this.length++
}


/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int)  {
    curtPrev := this.end.prev
    newNode := &Node{val, nil, nil}
    this.end.prev = newNode
    newNode.next = this.end
    curtPrev.next = newNode
    newNode.prev = curtPrev
    this.length++
}


/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int)  {
    if index > this.length {
        return
    }
    p := this.start
    for i := 0; i < index; i++ {
        p = p.next
    }
    // insert the new node.
    newNode := &Node{val, nil, nil}
    nextNode := p.next
    p.next = newNode
    newNode.prev = p
    newNode.next = nextNode
    nextNode.prev = newNode
    this.length++
}


/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int)  {
    if index >= this.length {
        return
    }
    p := this.start
    for i := 0; i < index; i++ {
        p = p.next
    }
    p.next = p.next.next
    p.next.prev = p
    this.length--
}