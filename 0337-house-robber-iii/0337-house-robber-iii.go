/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rob(root *TreeNode) int {
    
    result := recur(root)
    return Max(result[0], result[1])
}


func recur(root *TreeNode) []int{
    if root==nil{
        return []int{0,0}
    }
    
    l := recur(root.Left)
    r := recur(root.Right)
    
    withRoot := root.Val + l[1] + r[1]
    withOutRoot := Max(l[0], l[1]) + Max(r[0], r[1])

    return []int{withRoot, withOutRoot}
}


func Max(i, j int)int{
    if i>j{
        return i
    }
    return j
}