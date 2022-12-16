/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    
    var odds *ListNode = nil
    var evens *ListNode = nil
    var evenHead *ListNode = nil
    
    cur := head
    i := 0
    for cur != nil {
        i++
        next := cur.Next
        if i % 2 == 1 {
            if i > 1 {
                odds.Next = cur
            }
            odds = cur
            odds.Next = nil
        } else {
            if i == 2 {
                evenHead = cur
            } else {
                evens.Next = cur
            }
            evens = cur
            evens.Next = nil
        }
        cur = next
    }
    if odds != nil {
        odds.Next = evenHead
        return head
    }
    return evenHead
}