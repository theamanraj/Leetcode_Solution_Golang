/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertionSortList(head *ListNode) *ListNode {
    	pre := &ListNode{0, head}
	for head != nil && head.Next != nil {
		for ; head.Next != nil && head.Val <= head.Next.Val; head = head.Next {
		}
		if head.Next == nil {
			break
		}
		cur := head.Next
		next := cur.Next
		tmp := pre
		for ; cur.Val > tmp.Next.Val; tmp = tmp.Next {
		}
		head.Next = next
		cur.Next = tmp.Next
		tmp.Next = cur
	}
	return pre.Next
}