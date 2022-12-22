/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findMode(root *TreeNode) []int {
	_, _, _, T := f(root, 0, 0, 0, nil)
	return T
}

func f(node *TreeNode, preVal int, cnt int, mcnt int, T []int) (int, int, int, []int) {
	if node == nil {
		return preVal, cnt, mcnt, T
	}

	preVal, cnt, mcnt, T = f(node.Left, preVal, cnt, mcnt, T)
	if preVal != node.Val {
		cnt = 0
	}
	preVal = node.Val
	cnt++
	if cnt > mcnt {
        mcnt = cnt
		T = []int{node.Val}
	} else if cnt == mcnt {
		T = append(T, node.Val)
	}
	preVal, cnt, mcnt, T = f(node.Right, preVal, cnt, mcnt, T)

	return preVal, cnt, mcnt, T
}