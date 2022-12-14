/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
        left, right := &ListNode{0, nil}, &ListNode{0, nil}
        l, r := left, right
        node := head
        for node != nil {
                if node.Val < x {
                        l.Next = node
                        l = l.Next
                } else {
                        r.Next = node
                        r = r.Next
                }
                node = node.Next
        }
        r.Next = nil
        l.Next = right.Next
        return left.Next
}