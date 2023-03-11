/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func convertBST(root *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }
    recurse(root, 0)
    return root
}

func recurse(node *TreeNode, add int) int {
    if node.Right != nil {
        node.Val += recurse(node.Right, add)
    } else {
		//  If no right child, just add parent's value
        node.Val += add
    }
    if node.Left != nil {
        return recurse(node.Left, node.Val)
    }
    return node.Val
}