/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isCompleteTree(root *TreeNode) bool {
    if root==nil { return true}
    
    queue := []*TreeNode{root}    
    var isEnd bool
    
    for len(queue)>0{
        var nextDepth []*TreeNode
        for _, node := range queue{          
            if node==nil{
                isEnd=true
                continue
            }
            if isEnd { return false}
            nextDepth=append(nextDepth, node.Left)
            nextDepth=append(nextDepth, node.Right)
        }        
        queue = nextDepth        
    }
    return true    
}