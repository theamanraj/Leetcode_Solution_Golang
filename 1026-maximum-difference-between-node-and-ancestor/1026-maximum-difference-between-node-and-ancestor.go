/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func helper(node *TreeNode, minVal int, maxVal int, maxDiff *int) {
    if node == nil {
        return 
    }
    
    if abs(minVal - node.Val) > *maxDiff {
        *maxDiff = abs(minVal - node.Val)
    }
    
    if abs(maxVal - node.Val) > *maxDiff {
        *maxDiff = abs(maxVal - node.Val)
    }
    
    helper(node.Left, min(minVal, node.Val), max(maxVal, node.Val), maxDiff)
    helper(node.Right, min(minVal, node.Val), max(maxVal, node.Val), maxDiff)
}

func maxAncestorDiff(root *TreeNode) int {

    maxDiff := 0
    
    helper(root, root.Val, root.Val, &maxDiff)
    
    return maxDiff
}