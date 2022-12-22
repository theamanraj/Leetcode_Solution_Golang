/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findTilt(root *TreeNode) int {
    if root == nil {
        return 0
    }
    tiltAmt := 0
    helper(root, &tiltAmt) // helper function returns total value of root and calculate the tilt amount
    return tiltAmt
}

func helper(root *TreeNode, tiltAmt *int) int {
    if root == nil {
        return 0 // nil root's total value is 0, this is the recursion's exit point
    }
    totalLeft := helper(root.Left, tiltAmt) 
    totalRight := helper(root.Right, tiltAmt)
    *tiltAmt += abs(totalLeft - totalRight) //sum up tilt amount
    return root.Val + totalLeft + totalRight //return the current node's total value
}

func abs(a int) int {
    if a < 0 {
        return -a
    } else {
        return a
    }
}