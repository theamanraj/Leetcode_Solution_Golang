/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var data [][]int
func widthOfBinaryTree(root *TreeNode) int {
    data = [][]int{}
    Check(root, 0, 1)
    ans := 0
    tem := 0
    for _, v := range data {
        tem = 0
        tem = len(v)
        i := v[0]
        j := v[len(v)-1]
        
        tem = j-i+1
        if tem > ans {
            ans = tem
        }
    }
    return ans
}

func Check(root *TreeNode, xur int, l int) {
    if root == nil {
        return
    }
    if xur >= len(data) {
        data = append(data, []int{})
    }
    data[xur] = append(data[xur], l)

    if root.Left == nil {
    } else {
        if l == 0 {
            Check(root.Left, xur+1, 0)
        } else {
            Check(root.Left, xur+1, 2*l-1)
        }
    }
    

    if root.Right == nil {
    } else {
        if l == 0 {
            Check(root.Right, xur+1, 1)
        } else {
            Check(root.Right, xur+1, 2*l)
        }
    }
}