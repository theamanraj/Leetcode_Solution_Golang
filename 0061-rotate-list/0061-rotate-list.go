/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    
    curr := head
    l := 1
    
    for curr.Next != nil {
        l++
        curr = curr.Next
    }
    
    curr.Next = head
    
    k = k % l
    m := l-k-1
    
    for m > 0 {
        head = head.Next
        m--
    }
    
    newHead := head.Next
    head.Next = nil
    
    return newHead
}