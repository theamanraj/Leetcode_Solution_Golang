/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
    // Naively each level pushes the next level 
    // onto a stack that way the next level is called
    // using a lifo order (i.e the opposit of the previous level).
    if root == nil {
        return nil
    }
    
    res := make([][]int, 0)
    type TreeLevelNode struct {
        node *TreeNode
        level int
    }
    
    stack := make([]*TreeLevelNode,1)
    stack2 := make([]*TreeLevelNode,0)
    stack[0] = &TreeLevelNode{root,0}
    currLevel := 0;
    for len(stack) > 0 || len(stack2) > 0 {
        if currLevel % 2 == 0 {
            for len(stack) > 0 {
                nodeWithLevel := stack[len(stack)-1]
                stack = stack[:len(stack)-1]
                level := nodeWithLevel.level
                node := nodeWithLevel.node
                if len(res) < level + 1 {
                    res = append(res, []int{node.Val})
                } else {
                    res[level] = append(res[level], node.Val)
                }
                if node.Left != nil {
                   stack2 = append(stack2, &TreeLevelNode{node.Left,level+1})
                }
                if node.Right != nil {
                   stack2 = append(stack2, &TreeLevelNode{node.Right,level+1}) 
                }
            }
        } else {
            for len(stack2) > 0 {
                nodeWithLevel := stack2[len(stack2)-1]
                stack2 = stack2[:len(stack2)-1]
                level := nodeWithLevel.level
                node := nodeWithLevel.node
                if len(res) < level + 1 {
                    res = append(res, []int{node.Val})
                } else {
                    res[level] = append(res[level], node.Val)
                }
                if node.Right != nil {
                    stack = append(stack, &TreeLevelNode{node.Right,level+1})
                }
                if node.Left != nil {
                    stack = append(stack, &TreeLevelNode{node.Left,level+1}) 
                }
            }
        }
        currLevel += 1
    }

    return res
}