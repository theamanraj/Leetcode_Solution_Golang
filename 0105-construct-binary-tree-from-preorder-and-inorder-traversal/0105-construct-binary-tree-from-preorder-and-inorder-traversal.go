/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
    if len(preorder) == 0 || len(inorder) == 0 {
        return nil
    }
    var root = &TreeNode{Val: preorder[0]}
    var stack = []*TreeNode{root}
    var p, i = 1, 0
    for len(stack) > 0 {
        cur := stack[len(stack)-1]
        if inorder[i] != cur.Val {
            left := &TreeNode{Val: preorder[p]}
            p++
            cur.Left = left
            stack = append(stack, left)
            continue
        }
        i++
        stack = stack[:len(stack)-1]
        if len(stack) > 0 && stack[len(stack)-1].Val == inorder[i] {
            continue
        }
        if p >= len(preorder) {
            break
        }
        right := &TreeNode{Val: preorder[p]}
        p++
        cur.Right = right
        stack = append(stack, right)
    }
    return root
}