/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pruneTree(root *TreeNode) *TreeNode {
    contains := ContainsOne(root)
    if contains {
        return root
    }
    return nil
}

func ContainsOne(node *TreeNode) bool {
    if node == nil {
        return false
    }
    
    containsOne := (node.Val == 1)
    
    if ContainsOne(node.Left) {
       containsOne = true 
    } else {
        node.Left = nil
    }
    
    if ContainsOne(node.Right) {
        containsOne = true
    } else {
        node.Right = nil
    }
    
    return containsOne
}