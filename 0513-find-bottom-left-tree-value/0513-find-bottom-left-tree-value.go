/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findBottomLeftValue(root *TreeNode) int {
   	var maxHeight, res int
	var dfs func(root *TreeNode, curHeight int)
	dfs = func(root *TreeNode, curHeight int) {
		if root == nil {
			return
		}
		if root.Left == nil && root.Right == nil {
			if curHeight > maxHeight {
				res = root.Val
                maxHeight = curHeight
			}
		} else {
			dfs(root.Left, curHeight+1)
			dfs(root.Right, curHeight+1)
		}
	}
	dfs(root, 1)
	return res 
}