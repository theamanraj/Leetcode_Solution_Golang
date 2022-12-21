/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func levelOrder(root *Node) [][]int {
    
    if root == nil {
        return [][]int{}
    }
    
    result := [][]int{}
    
    queue := []*Node{}
    
    queue = append(queue, root)
    
    for len(queue) != 0 {
        n := len(queue)
        temp := []int{}
        for i := 0; i < n; i++ {
            node := queue[0]
            queue = queue[1:]
            
            for _, child := range node.Children {
                queue = append(queue, child)
            }
            
            temp = append(temp, node.Val)
        }
        
        result = append(result, temp)
        
    }
    
    return result
}