/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, m int, n int) *ListNode {
    dummy := &ListNode{}
    dummy.Next = head
    head = dummy
    for i := 0; i < m - 1; i++ { head = head.Next }
    var prev, cur *ListNode = nil, head.Next
    for i := 0; i <= n - m; i++ {
        next := cur.Next
        cur.Next = prev
        prev = cur
        cur = next
    }
    head.Next.Next = cur
    head.Next = prev
    return dummy.Next
}