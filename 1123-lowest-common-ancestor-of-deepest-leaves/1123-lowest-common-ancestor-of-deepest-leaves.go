/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	node, _:= helper(root, 1)

	return node
}

func helper(root *TreeNode, lvl int) (*TreeNode, int) {
	if root == nil {
		return nil, lvl-1
	}

	left, ll := helper(root.Left, lvl+1)
	right, rl := helper(root.Right, lvl+1)

	if ll == rl {
		return root, ll
	}
	
	if ll > rl {
		return left, ll
	}

	return right, rl
}