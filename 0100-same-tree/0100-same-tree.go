/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
    if p == nil || q == nil {
        return p == q
    }
    cur1, cur2, s1, s2 := p, q, []*TreeNode{}, []*TreeNode{}
    for cur1 != nil || len(s1) != 0 {
        for cur1 != nil {
            if cur2 == nil || cur1.Val != cur2.Val {
                return false
            }
            s1, s2 = append(s1, cur1), append(s2, cur2)
            cur1, cur2 = cur1.Left, cur2.Left
        }
        if cur2 != nil {
            return false
        }
        cur1, cur2 = s1[len(s1) - 1].Right, s2[len(s2) - 1].Right
        s1, s2 = s1[:len(s1) - 1], s2[:len(s2) - 1]
        if cur1 == nil && cur2 != nil {
            return false
        }
    }
    return true
}