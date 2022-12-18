/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var paths [] string

func dfs(root * TreeNode, path string) {
    if(root.Left == nil && root.Right == nil) {
        paths = append(paths, path + strconv.Itoa(root.Val))
		return
    }
    if(root.Left != nil) {
        dfs(root.Left, path + strconv.Itoa(root.Val) + "->")     
    }
    if(root.Right != nil) {
        dfs(root.Right, path + strconv.Itoa(root.Val) + "->")
    }
}
func binaryTreePaths(root *TreeNode) []string {
    paths = make([]string, 0)
    if(root == nil) {
        return paths
    }
    dfs(root, "")
    return paths
}