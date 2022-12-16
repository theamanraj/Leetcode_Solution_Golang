/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type dfsNode struct {
    *TreeNode
    sum int
}

func sumNumbers(root *TreeNode) int {
    if root == nil {
        return 0
    }
    sum := 0
    for stack := []dfsNode{dfsNode{root, 0}}; len(stack) != 0; {
        current := stack[len(stack) - 1]
        stack = stack[:len(stack) - 1]
        if current.Left == nil && current.Right == nil {
            sum += current.sum * 10 + current.Val
            continue
        }
        if current.Right != nil {
            stack = append(stack, dfsNode{current.Right, current.sum * 10 + current.Val})
        }
        if current.Left != nil {
            stack = append(stack, dfsNode{current.Left, current.sum * 10 + current.Val})
        }            
    }                
    return sum
}
