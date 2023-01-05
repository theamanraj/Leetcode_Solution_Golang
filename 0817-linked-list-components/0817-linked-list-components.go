/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func numComponents(head *ListNode, G []int) int {
	m := make(map[int]struct{}, len(G)) // len(G) is important: 75%->96%
	for _, v := range G {
		m[v] = struct{}{}
	}
	c := 0
	for head != nil {
		if _, ok := m[head.Val]; ok {
			c++
			for head != nil {
				if _, ok := m[head.Val]; !ok {
					break
				}
				head = head.Next
			}
		} else {
			head = head.Next
		}
	}
	return c
}