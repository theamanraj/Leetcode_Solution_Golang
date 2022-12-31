/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	d1, d2, d3 := &ListNode{}, &ListNode{}, &ListNode{}
	d1.Next, d2.Next = l1, l2
	rl1, rl2 := reverseList(d1), reverseList(d2)
	p1, p2, p3 := rl1, rl2, d3
	var num, addOne, curVal int
	for p1 != nil || p2 != nil {
		if p1 != nil && p2 != nil {
			curVal = p1.Val + p2.Val
			p1 = p1.Next
			p2 = p2.Next
		} else if p1 != nil {
			curVal = p1.Val
			p1 = p1.Next
		} else {
			curVal = p2.Val
			p2 = p2.Next
		}
		num = (curVal + addOne) % 10
		addOne = (curVal + addOne) / 10
		nn := &ListNode{Val: num}
		p3.Next = nn
		p3 = nn
	}
    if addOne > 0 {
		nn := &ListNode{Val: addOne}
		p3.Next = nn
	}
	res := reverseList(d3)
	return res
}

func reverseList(dummy *ListNode) *ListNode {
	var prev *ListNode
	p := dummy.Next
	for p.Next != nil {
		next := p.Next
		p.Next = prev
		prev = p
		p = next
	}
	p.Next = prev
	return p
}
