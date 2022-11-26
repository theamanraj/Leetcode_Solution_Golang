func detectCycle(head *ListNode) *ListNode {
    slower := head
    faster := head
    for faster != nil && faster.Next != nil {
        slower = slower.Next
        faster = faster.Next.Next
        if slower == faster {
            faster = head
            for faster != slower { // the list may be a circular list
                faster = faster.Next
                slower = slower.Next
            }
            return faster
        }
    }
    
    return nil
}