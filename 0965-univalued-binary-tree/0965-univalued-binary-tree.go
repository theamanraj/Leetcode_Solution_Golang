/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type Stack []*TreeNode
func (s *Stack) Push(t *TreeNode) {
    *s = append(*s, t)
}
func (s *Stack) Pop() *TreeNode{
    l := len(*s)
    val := (*s)[l-1]
    *s = (*s)[:l-1]
    return val
}
func (s *Stack) isEmpty() bool {
    return len(*s) < 1
}

func isUnivalTree(root *TreeNode) bool {
    
    stack := make(Stack,0)
    stack.Push(root)
    univalue := root.Val
    
    for !stack.isEmpty() {
        node := stack.Pop()
        
        if node != nil {
            
            //not a univalued tree
            if node.Val != univalue {
                return false
            }
            
            stack.Push(node.Right)
            stack.Push(node.Left)
        }
    }
    
    return true
}