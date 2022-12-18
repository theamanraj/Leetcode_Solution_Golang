/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func evaluateTree(root *TreeNode) bool {
    if root.Left != nil { // non-leaf
        a := evaluateTree(root.Left)
        b := evaluateTree(root.Right)
        if root.Val == 2 {
            return a || b
        } else if root.Val == 3{
            return a && b
        }
    }
    // leaf
    if root.Val == 0 {
        return false
    } else if root.Val == 1 {
        return true
    }
    return false
}