func flattenHelper(root *Node) *Node {
    curr := root
    prev := root
    if curr == nil {
        return nil
    }
    for ;curr != nil;curr = curr.Next {
        if curr.Child != nil {
            tmp := curr.Next
            // Link the child with curr
            curr.Next = curr.Child
            curr.Child.Prev = curr
            
            // Recursively flatten and get the last non-empty node
            node := flattenHelper(curr.Child)
            // Link the last node with the actual next node
            if tmp != nil {
                tmp.Prev = node
            }
            node.Next = tmp
            curr.Child = nil
        }
        prev = curr
    }
    return prev
}
func flatten(root *Node) *Node {
    dummyNode := new(Node)
    dummyNode.Next = root
    flattenHelper(root)
    return dummyNode.Next
}