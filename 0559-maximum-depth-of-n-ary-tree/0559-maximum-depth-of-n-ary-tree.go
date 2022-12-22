/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func maxDepth(root *Node) int {
  if root == nil {
    return 0
  }
  
  return dfs(root, 1, 0)
}

func dfs(root *Node, level, count int) int {
  if root == nil {
    return count
  }
  
  count = int(math.Max(float64(count), float64(level)))
  for _, node := range root.Children {
    count = dfs(node, level+1, count)
  }
  
  return count
}

