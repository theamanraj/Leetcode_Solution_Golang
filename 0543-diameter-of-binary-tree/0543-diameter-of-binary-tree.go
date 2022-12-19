/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var height int
func diameterOfBinaryTree(root *TreeNode) int {
    height = 0
    _ = find(root)
    return height
}


func find(root *TreeNode) int {
    if root == nil {
        return 0
    }
    if root.Left == nil && root.Right == nil {
        return 1
    }    
    left := find(root.Left)
    right := find(root.Right)
    height = max(height, left + right)
    return max(left, right) + 1
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}