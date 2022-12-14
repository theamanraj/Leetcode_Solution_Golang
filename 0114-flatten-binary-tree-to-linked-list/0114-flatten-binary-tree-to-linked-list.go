/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode)  {
    root = helper(root)
}
func helper(root *TreeNode) *TreeNode{
    if root == nil || (root.Left == nil && root.Right==nil) {
        return root
    }
    left := root.Left
    if root.Left != nil {
        left = helper(root.Left)
    }
    right := root.Right
    if root.Right != nil {
        right = helper(root.Right)
    }
    if left!=nil {
        temp := left
        for left.Right!=nil {
            left = left.Right
        }
        left.Right=right        
        root.Right = temp
        root.Left =nil
    }
    return root
}