/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedListToBST(head *ListNode) *TreeNode {
    if head == nil { return nil }
    dummy := new(ListNode)
    head, dummy.Next = dummy, head
    fast, slow := head.Next, head.Next
    for fast != nil && fast.Next != nil {
        dummy = dummy.Next
        slow = slow.Next
        fast = fast.Next.Next
    }
    dummy.Next = nil
    return &TreeNode {
        slow.Val,
        sortedListToBST(head.Next),
        sortedListToBST(slow.Next),
    }
}