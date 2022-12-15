/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func generateTrees(n int) []*TreeNode {
    if n == 0 {
        return []*TreeNode{}
    } else if n == 1 {
        return []*TreeNode{
            &TreeNode{Val: 1},
        }
    }
    trees := []*TreeNode{}
    for _, tree := range generateTrees(n-1) {
        temp := tree
        root := &TreeNode{Val: n, Left: cloneTree(temp)}
        trees = append(trees, root)
        for temp != nil {
            temp.Right = &TreeNode{Val: n, Left: temp.Right}
            trees = append(trees, cloneTree(tree))
            temp.Right = temp.Right.Left
            temp = temp.Right
        }
    }
    return trees
}

func cloneTree(t *TreeNode) *TreeNode {
	if t == nil {
		return nil
	}
    return &TreeNode{
        Val: t.Val,
        Left: cloneTree(t.Left),
        Right: cloneTree(t.Right),
    }
}