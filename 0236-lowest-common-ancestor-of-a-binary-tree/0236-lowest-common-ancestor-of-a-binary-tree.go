/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
     if root == nil {
         return nil
     } else if root == p {
         return p
     } else if root == q {
         return q
     }
     
     lc := lowestCommonAncestor(root.Left, p, q)
     rc := lowestCommonAncestor(root.Right, p, q)
     
     if lc != nil && rc != nil {
         return root
     } else if lc != nil {
         return lc
     } else {
         return rc
     }
}