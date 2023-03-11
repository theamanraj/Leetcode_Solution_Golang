/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flipEquiv(root1 *TreeNode, root2 *TreeNode) bool {
    if root1 != nil && root2 != nil {
        if root1.Val != root2.Val {
            return false
        }
        if flipEquiv(root1.Right, root2.Left) && flipEquiv(root1.Left, root2.Right){
            return true
        }
        if flipEquiv(root1.Left, root2.Left) && flipEquiv(root1.Right, root2.Right) {
            return true
        }
    } else if root1 == nil && root2 == nil {
        return true
    }
    return false
}