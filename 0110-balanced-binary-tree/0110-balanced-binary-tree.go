/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
    if root == nil {
        return true
    }
    return helper(root) != -1
}

func helper(root *TreeNode) int {
    if root == nil {
        return 0
    }
    left := helper(root.Left)
    right := helper(root.Right)
    if (left == -1 || right == -1 || abs(left - right) > 1) {
        return -1
    }
    return max(left, right) + 1
}

func abs(i int) int {
    if i < 0 {
        return i * -1
    }
    return i
}

func max(a int, b int) int {
    if a > b {
        return a
    }
    return b
}