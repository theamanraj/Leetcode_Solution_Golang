/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getParent(root *TreeNode, x int) (*TreeNode, int) {
    if root == nil { return nil, -1}
    if root.Val == x { return nil, 0}
    if root.Left != nil && root.Left.Val == x { return root, 1}
    if root.Right != nil && root.Right.Val == x { return  root, 1}
    lp, ld := getParent(root.Left, x)
    rp, rd := getParent(root.Right, x)
    if lp == nil { return rp , rd + 1}
    return lp, ld + 1
}

func isCousins(root *TreeNode, x int, y int) bool {
    xp, xd := getParent(root, x)
    yp, yd := getParent(root, y)
    if xp != yp &&  xd == yd { return true }
    return false
}