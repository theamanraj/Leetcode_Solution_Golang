/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	var res []int
	processFunc := func(val int) {
		res = append(res, val)
	}
	Helper(head, processFunc)
	sort.Ints(res)
	dummy := &ListNode{}
	cur := dummy
	for _, v := range res {
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
	}
	return dummy.Next
}

func Helper(head *ListNode, processFunc func(val int)) {
	if head == nil {
		return
	}
	processFunc(head.Val)
	Helper(head.Next, processFunc)
}