/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getMiddleNode(node *ListNode) *ListNode {
    slow, fast := node, node

    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }

    return slow
}

func reverseList(node *ListNode, prev *ListNode) *ListNode {
    if node == nil {
        return prev
    }

    nxt := node.Next
    node.Next = prev
    return reverseList(nxt, node)
}

func removeLastNode(node *ListNode) *ListNode {
    if node == nil {
        return node
    }

    curr := node
    for curr.Next != nil && curr.Next.Next != nil {
        curr = curr.Next
    }

    curr.Next = nil
    return node
}

func reorderList(head *ListNode)  {
    if head == nil {
        return
    }
    
    midNode := getMiddleNode(head)
    reversed := reverseList(midNode, nil)
    dummy := removeLastNode(head)
    ptr1, ptr2 := dummy, reversed

    for ptr1 != nil {
        nxt1, nxt2 := ptr1.Next, ptr2.Next
        ptr1.Next = ptr2
        ptr2.Next = nxt1
        ptr1, ptr2 = nxt1, nxt2
    }

    if ptr2 != nil {
        curr := dummy
        for curr.Next != nil {
            curr = curr.Next
        }
        curr.Next = ptr2
    }

}