/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
    var ans int
    var dfs func(*TreeNode)
    
    dfs = func(node *TreeNode) {
        if node == nil {
            return
        }
        
        dfs(node.Left)
        k--
        if k == 0 {
           ans = node.Val    
            return
        }
        dfs(node.Right)
    }
    
    dfs(root)
    return ans
}