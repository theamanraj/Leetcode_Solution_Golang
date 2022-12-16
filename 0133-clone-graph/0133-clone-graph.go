/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
    if node == nil {
        return nil
    }
    visited := make(map[int]*Node)
    return walk(node, visited)
}

func walk(node *Node, visited map[int]*Node) *Node {
    if n, ok := visited[node.Val]; ok {
        return n
    }
    n := &Node{Val: node.Val, Neighbors: make([]*Node, len(node.Neighbors))}
    visited[node.Val] = n
    for i, _ := range node.Neighbors {
        n.Neighbors[i] = walk(node.Neighbors[i], visited)
    }
    return n
}