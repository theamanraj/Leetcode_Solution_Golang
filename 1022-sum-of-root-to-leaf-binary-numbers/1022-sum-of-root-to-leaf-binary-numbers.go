/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumRootToLeaf(root *TreeNode) int {
    
    return recursion(root, 0, 0)
}

func recursion(root *TreeNode, nxt int, sum int) int {
    
    nxt = nxt << 1 + root.Val

    if root.Left == nil && root.Right == nil {
       sum += nxt
       return sum
    }
                                                 
    if root.Left != nil {
       sum = recursion(root.Left, nxt, sum)
    }

    if root.Right != nil {
        sum = recursion(root.Right, nxt, sum)
    }

    return sum
}