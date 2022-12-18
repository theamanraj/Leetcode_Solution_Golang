/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// Using BFS due to comment: https://leetcode.com/problems/minimum-depth-of-binary-tree/discuss/36045/My-4-Line-java-solution/167883
func minDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    
    q := []*TreeNode{root}
    lvl := 1
    
    for len(q) > 0 {
        size := len(q)
        
		// Inner loop is required to check Left node and Right node (if one of them is a leaf) before increasing the level.
        for i := 0; i < size; i++ {
            node := q[i]

            if node == nil {
                continue
            }
            
            if node.Left == nil && node.Right == nil {
                return lvl
            }
            
            q = append(q, node.Left, node.Right)
        }
        // Resize the slice to avoid adding the same elements.
        q = q[size:]
        
        lvl++
    }
    
    return lvl
}