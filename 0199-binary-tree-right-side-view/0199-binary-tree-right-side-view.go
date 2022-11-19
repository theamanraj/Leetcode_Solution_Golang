/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
    arr:= levelOrderBottom(root)
    res:=[]int{}
    for _,e:=range arr {
        res = append(res, e[len(e)-1])
    }
    return res
}

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	ans := [][]int{}
	dfs(root, 0, &ans)
	return ans
}

func dfs(node *TreeNode, deep int, ans *[][]int) {
	if len(*ans) <= deep {
		*ans = append( *ans ,[]int{})
	}
	(*ans)[deep] = append((*ans)[deep], node.Val)
	if node.Left != nil {
		dfs(node.Left, deep+1, ans)
	}
	if node.Right != nil {
		dfs(node.Right, deep+1, ans)
	}
}