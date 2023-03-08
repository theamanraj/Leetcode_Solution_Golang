/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func bstToGst(root *TreeNode) *TreeNode {
	var dfs func(node *TreeNode, greaterSum int) int
	dfs = func(node *TreeNode, greaterSum int) int {
		if node == nil {
			return 0
		}
		r := dfs(node.Right, greaterSum)
		tmp := node.Val
		node.Val += r + greaterSum
		return r + tmp + dfs(node.Left, node.Val)
	}
	dfs(root, 0)
	return root
}