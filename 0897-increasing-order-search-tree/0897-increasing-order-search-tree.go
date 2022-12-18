/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func increasingBST(root *TreeNode) *TreeNode {
    dummy = &TreeNode{}
    now = dummy
    
    visit(root)
    return dummy.Right
}

var (
    dummy = &TreeNode{}
    now = dummy
)
    

func visit(node *TreeNode){
    if node == nil{
        return
    }
    
    if node.Left != nil{
        visit(node.Left)
    }
    
    now.Right = &TreeNode{Val: node.Val}
    now = now.Right
    
    if node.Right != nil{
        visit(node.Right)
    }
}