/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func largestValues(root *TreeNode) []int {
	res := []int{}
	largestValuesHelper(root, &res, 0)
	return res
}

func largestValuesHelper(node *TreeNode, res *[]int, level int) {
	if node == nil {
		return
	}

	if len(*res) == level {
		*res = append(*res, node.Val)
	} else {
		(*res)[level] = max((*res)[level], node.Val)
	}

	largestValuesHelper(node.Left, res, level+1)
	largestValuesHelper(node.Right, res, level+1)
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}