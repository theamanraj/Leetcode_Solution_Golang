/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func recoverTree(root *TreeNode) {
	var swap1 *TreeNode
	var swap2 *TreeNode

	cur := root
	stack := []*TreeNode{}

	var prev *TreeNode
	for cur != nil || len(stack) > 0 {
		if cur == nil {
			cur = stack[len(stack)-1]
			if prev != nil && prev.Val > cur.Val {
				swap2 = cur
				if swap1 == nil {
					swap1 = prev
				}
			}

			prev = cur
			cur = cur.Right
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, cur)
			cur = cur.Left
		}
	}

	swap1.Val, swap2.Val = swap2.Val, swap1.Val
}