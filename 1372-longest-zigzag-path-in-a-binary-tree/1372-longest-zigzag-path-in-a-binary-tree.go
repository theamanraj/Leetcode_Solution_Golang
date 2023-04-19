/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func longestZigZag(root *TreeNode) int {
    current_max := 0

    if root.Left != nil {
        dfs(root.Left, 1, false,&current_max)
    }   
    if root.Right != nil {
        dfs(root.Right, 1, true,&current_max)
    }
    return current_max 
}


func dfs(node *TreeNode, step int, pre_right bool, current_max *int) {
    if step > *current_max {
        *current_max = step
    }

    if pre_right {
        if node.Left != nil {
            dfs(node.Left,step+1,false,current_max)
        }
        if node.Right != nil {
            dfs(node.Right,1,true,current_max)
        }
    }else{
        if node.Left != nil {
            dfs(node.Left,1,false,current_max)
        }
        if node.Right != nil {
            dfs(node.Right,step+1,true,current_max)
        }
    }
}
