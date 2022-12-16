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
func isSubPath(head *ListNode, root *TreeNode) bool {
    queue := []*TreeNode{root}
    for len(queue) != 0 {
        cur := queue[0]
        if len(queue) == 1{
            queue = []*TreeNode{}
        }else{
            queue = queue[1:]
        }
        if cur.Left != nil {
            queue = append(queue, cur.Left)
        }
        if cur.Right != nil {
            queue = append(queue, cur.Right)
        }
        if check(head, cur) {
            return true
        }
    }
    return false
}

func check(head *ListNode, root *TreeNode) bool {
    if head == nil {
        return true
    }
    if root == nil {
        return false
    }
    
    if head.Val == root.Val {
        return check(head.Next, root.Left) || check(head.Next, root.Right)
    }
    return false
}