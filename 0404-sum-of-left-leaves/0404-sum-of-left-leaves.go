/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumOfLeftLeaves(root *TreeNode) int {
	res := 0
	sumOfLeftLeavesHelper(root, false, &res)
	return res
}
func sumOfLeftLeavesHelper(root *TreeNode, isLeftNode bool, res *int) {
	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil && isLeftNode {
		*res += root.Val
	}

	sumOfLeftLeavesHelper(root.Left, true, res)
	sumOfLeftLeavesHelper(root.Right, false, res)
}