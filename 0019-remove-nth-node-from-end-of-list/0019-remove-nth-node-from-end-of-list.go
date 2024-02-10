func removeNthFromEnd(head *ListNode, n int) *ListNode {
    if head == nil || n <= 0 {return head}
    preHeader := &ListNode{Next: head}
    left, right := preHeader, head

    for i := 0; right != nil && i < n; i++ {
        right = right.Next
    }
    for right != nil {
        left, right = left.Next, right.Next
    }
    left.Next = left.Next.Next
    return preHeader.Next
}